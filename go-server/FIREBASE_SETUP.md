# Firebase Setup Guide

## Quick Fix for "Firebase not initialized" Error

You're seeing this error because Firebase credentials are not properly configured. Here are the solutions:

### Option 1: Use Firebase Service Account Key File (Recommended)

1. **Get your Firebase service account key:**
   - Go to [Firebase Console](https://console.firebase.google.com/)
   - Select your project
   - Go to Project Settings (⚙️ icon) → Service Accounts
   - Click "Generate New Private Key"
   - Save the JSON file as `firebase-key.json` in your project root

2. **Update your `.env` file:**
   ```env
   FIREBASE_SERVICE_ACCOUNT_KEY=./firebase-key.json
   ```

3. **Restart the server**

### Option 2: Use Individual Environment Variables

If you can't use a service account key file, add these to your `.env`:

```env
# Firebase Web App Config (for client-side)
FIREBASE_API_KEY=your-api-key
FIREBASE_AUTH_DOMAIN=your-project.firebaseapp.com
FIREBASE_PROJECT_ID=your-project-id
FIREBASE_STORAGE_BUCKET=your-project.appspot.com
FIREBASE_MESSAGING_SENDER_ID=your-sender-id
FIREBASE_APP_ID=your-app-id

# Firebase Admin SDK (for server-side)
FIREBASE_PROJECT_ID=your-project-id
FIREBASE_CLIENT_EMAIL=firebase-adminsdk-xxxxx@your-project.iam.gserviceaccount.com
FIREBASE_PRIVATE_KEY="-----BEGIN PRIVATE KEY-----\nYour-Private-Key-Here\n-----END PRIVATE KEY-----\n"
```

### Option 3: Temporarily Disable Firebase (For Testing)

If you just want to test without authentication:

1. Comment out Firebase initialization in `cmd/server/main.go`:
   ```go
   // if err := config.InitializeFirebase(); err != nil {
   //     utils.Warn("Firebase initialization failed: %v", err)
   // }
   ```

2. Comment out the `RequireAuth` middleware in `internal/routes/routes.go`:
   ```go
   // adminGroup.Use(middleware.RequireAuth)
   ```

## Current Error Explanation

The error message shows:
```
⚠️  Firebase service account key appears to be JSON content, not a file path
⚠️  Firebase Auth client creation failed: google: could not find default credentials
```

This means:
- The `FIREBASE_SERVICE_ACCOUNT_KEY` in your `.env` is set to JSON content (not a file path)
- Firebase can't find valid credentials to initialize

## Recommended Solution

**Create `firebase-key.json` file:**

1. Download your Firebase service account key from Firebase Console
2. Save it as `firebase-key.json` in the project root
3. Update `.env`:
   ```env
   FIREBASE_SERVICE_ACCOUNT_KEY=./firebase-key.json
   ```
4. Restart the server

The admin panel will then work with proper authentication!

## Restrict Admin Access

Firebase authentication proves who a user is; it does not make that user an
administrator. Add every permitted admin email to `ADMIN_EMAILS` in your
environment configuration:

```env
ADMIN_EMAILS=owner@example.com,editor@example.com
```

Only Firebase accounts with a verified email address in this allowlist can
create an admin session. If this setting is empty, all admin logins are denied.
