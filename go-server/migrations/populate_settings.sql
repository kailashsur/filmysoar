-- Populate default settings from Node.js project
-- Run this in your PostgreSQL database

-- Insert default settings
INSERT INTO settings ("key", "value", "description", "createdAt", "updatedAt")
VALUES 
    ('downloadRedirectUrl', '', 'Redirect URL for download links', NOW(), NOW()),
    ('googleTagManagerHead', '', 'Google Tag Manager code for <head> section', NOW(), NOW()),
    ('googleTagManagerBody', '', 'Google Tag Manager noscript code for <body> section', NOW(), NOW()),
    ('googleAnalytics', '', 'Google Analytics (gtag.js) code', NOW(), NOW()),
    ('googleSearchConsole', '', 'Google Search Console verification meta tag', NOW(), NOW()),
    ('adsenseCode', '', 'Google AdSense code', NOW(), NOW()),
    ('adsteraCode', '', 'Adstera ad code', NOW(), NOW()),
    ('siteUrl', 'https://filmyfly.work', 'Site URL (used in meta tags, Open Graph, Twitter Card, etc.)', NOW(), NOW())
ON CONFLICT ("key") DO UPDATE SET
    "description" = EXCLUDED."description",
    "updatedAt" = NOW();

-- Insert default Astro settings
INSERT INTO astro_settings ("key", "value", "description", "createdAt", "updatedAt")
VALUES 
    ('downloadRedirectUrl', '', 'Redirect URL for download links (Astro)', NOW(), NOW()),
    ('siteTitle', 'FilmyFly', 'Site title for Astro frontend', NOW(), NOW()),
    ('siteDescription', 'Download Latest Movies', 'Site description for Astro frontend', NOW(), NOW()),
    ('apiBaseUrl', 'http://localhost:3000', 'API base URL for Astro frontend', NOW(), NOW())
ON CONFLICT ("key") DO UPDATE SET
    "description" = EXCLUDED."description",
    "updatedAt" = NOW();

-- Verify the inserts
SELECT 'Settings:' as table_name, COUNT(*) as count FROM settings
UNION ALL
SELECT 'Astro Settings:', COUNT(*) FROM astro_settings;
