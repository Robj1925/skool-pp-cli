const { chromium } = require('playwright');
const path = require('path');

(async () => {
  // Use a relative path so it works for any user
  const userDataDir = path.join(__dirname, '..', 'skoolbot_profile');
  console.log(`Launching browser with dedicated profile: ${userDataDir}`);
  let browser, page;
  try {
    browser = await chromium.launchPersistentContext(userDataDir, {
      headless: false,
    });
    page = await browser.newPage();
  } catch (err) {
    console.error(`FATAL: Failed to launch Playwright browser. Error: ${err.message}`);
    process.exit(1);
  }

  // CHECK IMMEDIATELY: Are we already logged in?
  const initialCookies = await browser.cookies("https://www.skool.com");
  if (initialCookies.some(c => c.name === 'auth_token')) {
    console.log("Existing session found in profile! Capturing...");
    const cookieString = initialCookies.map(c => `${c.name}=${c.value}`).join('; ');
    process.stdout.write(`COOKIE_RESULT=${cookieString}\n`, async () => {
      await browser.close();
      process.exit(0);
    });
    return;
  }

  console.log("No existing session found. Navigating to Skool...");
  await page.goto("https://www.skool.com/chat");

  const checkInterval = setInterval(async () => {
    try {
      const cookies = await browser.cookies("https://www.skool.com");
      if (cookies.some(c => c.name === 'auth_token')) {
        const cookieString = cookies.map(c => `${c.name}=${c.value}`).join('; ');
        clearInterval(checkInterval);
        process.stdout.write(`COOKIE_RESULT=${cookieString}\n`, async () => {
          await browser.close();
          process.exit(0);
        });
      }
    } catch (err) {
      // Ignore execution context errors during navigation
    }
  }, 2000);

  try {
    await page.waitForTimeout(600000); // Wait up to 10 mins
    console.error("Timeout: Could not detect auth_token in cookies after 10 minutes.");
    clearInterval(checkInterval);
    await browser.close();
    process.exit(1);
  } catch (err) {
    console.error(`Browser session interrupted: ${err.message}`);
    clearInterval(checkInterval);
    try { await browser.close(); } catch (e) {}
    process.exit(1);
  }
})();
