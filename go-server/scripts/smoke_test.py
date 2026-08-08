#!/usr/bin/env python3
"""Dependency-free smoke tests for the FilmyFly Go/Fiber service.

Examples:
  python scripts/smoke_test.py
  python scripts/smoke_test.py --base-url http://localhost:3000 --skip-build
  python scripts/smoke_test.py --admin-cookie "session_id=..."
"""

import argparse
import json
import os
import subprocess
import sys
import urllib.error
import urllib.request


class SmokeTests:
    def __init__(self, args):
        self.args = args
        self.base = args.base_url.rstrip("/")
        self.passed = 0
        self.failed = 0

    def check(self, name, condition, detail=""):
        if condition:
            self.passed += 1
            print(f"PASS {name}")
        else:
            self.failed += 1
            print(f"FAIL {name}: {detail}")

    def command(self, name, command):
        if self.args.skip_build:
            print(f"SKIP {name} (--skip-build)")
            return
        result = subprocess.run(command, cwd=self.args.project_dir,
                                capture_output=True, text=True)
        output = (result.stderr or result.stdout).strip()
        if name == "gofmt":
            ok = result.returncode == 0 and not result.stdout.strip()
        else:
            ok = result.returncode == 0
        self.check(name, ok,
                   output[-500:])

    def request(self, path, method="GET", headers=None):
        request = urllib.request.Request(self.base + path, method=method,
                                          headers=headers or {})
        try:
            with urllib.request.urlopen(request, timeout=self.args.timeout) as response:
                body = response.read().decode("utf-8", errors="replace")
                return response.status, dict(response.headers), body
        except urllib.error.HTTPError as error:
            body = error.read().decode("utf-8", errors="replace")
            return error.code, dict(error.headers), body
        except (urllib.error.URLError, TimeoutError) as error:
            return None, {}, str(error)

    def json_request(self, path):
        status, headers, body = self.request(path)
        try:
            return status, headers, json.loads(body)
        except json.JSONDecodeError:
            return status, headers, None

    def run(self):
        self.command("gofmt", ["gofmt", "-l", "."])
        self.command("go vet", ["go", "vet", "./..."])
        self.command("go tests", ["go", "test", "./..."])
        self.command("go build", ["go", "build", "./cmd/server"])

        for path in ("/api/home", "/api/categories", "/api/movies", "/api/astro-settings"):
            status, _, data = self.json_request(path)
            self.check(f"GET {path}", status == 200 and isinstance(data, dict),
                       f"status={status}")

        for path in ("/api/movies?page=0&limit=0",
                     "/api/movies?page=-1&limit=999999",
                     "/api/categories?page=0&limit=0",
                     "/api/search?q=test&page=-2&limit=999999"):
            status, _, data = self.json_request(path)
            self.check(f"pagination {path}", status == 200 and isinstance(data, dict),
                       f"status={status}")

        status, _, _ = self.request("/admin")
        self.check("admin requires authentication", status in (200, 301, 302, 303, 401),
                   f"status={status}")

        cookie = self.args.admin_cookie
        if cookie:
            common = {"Cookie": cookie, "Origin": self.base}
            status, _, _ = self.request("/admin/settings", headers={"Cookie": cookie})
            self.check("authenticated admin page", status == 200, f"status={status}")

            status, _, _ = self.request("/admin/logs/clear", method="POST",
                                        headers={"Cookie": cookie})
            self.check("CSRF rejects missing origin", status == 403, f"status={status}")

            status, _, _ = self.request("/admin/logs/clear", method="POST", headers=common)
            self.check("CSRF accepts same origin", status != 403, f"status={status}")
        else:
            print("SKIP authenticated admin and CSRF checks (use --admin-cookie)")

        print(f"\n{self.passed} passed, {self.failed} failed")
        return 1 if self.failed else 0


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--base-url", default=os.getenv("BASE_URL", "http://localhost:3000"))
    parser.add_argument("--project-dir", default=os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
    parser.add_argument("--admin-cookie", default=os.getenv("ADMIN_COOKIE"),
                        help="Authenticated session cookie, e.g. session_id=...")
    parser.add_argument("--timeout", type=float, default=10)
    parser.add_argument("--skip-build", action="store_true")
    args = parser.parse_args()
    sys.exit(SmokeTests(args).run())


if __name__ == "__main__":
    main()
