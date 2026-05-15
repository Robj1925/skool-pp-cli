const { chromium } = require('playwright');
const path = require('path');

(async () => {
  // Use a relative path so it works for any user
  const userDataDir = path.join(__dirname, '..', 'skoolbot_profile');
  console.log(`Launching browser with dedicated profile: ${userDataDir}`);
  
  const browser = await chromium.launchPersistentContext(userDataDir, {
    headless: false,
  });
  
  const page = await browser.newPage();

  // CHECK IMMEDIATELY: Are we already logged in?
  const initialCookies = await browser.cookies();
  if (initialCookies.some(c => c.name === 'auth_token')) {
    console.log("Existing session found in profile! Capturing...");
    const cookieString = initialCookies.map(c => `${c.name}=${c.value}`).join('; ');
    console.log(`COOKIE_RESULT=${cookieString}`);
    await browser.close();
    process.exit(0);
  }

  console.log("No existing session found. Navigating to Skool...");
  await page.goto("https://www.skool.com/chat");

  const checkInterval = setInterval(async () => {
    console.log(`Current URL: ${page.url()}`);
    const cookies = await browser.cookies();
    if (cookies.some(c => c.name === 'auth_token')) {
      console.log("Auth token detected in cookies!");
      const cookieString = cookies.map(c => `${c.name}=${c.value}`).join('; ');
      console.log(`COOKIE_RESULT=${cookieString}`);
      clearInterval(checkInterval);
      await browser.close();
      process.exit(0);
    }
  }, 2000);

  try {
    await page.waitForTimeout(120000); // Wait up to 2 mins for the interval check to succeed
    console.error("Timeout: Could not detect auth_token in cookies after 120 seconds.");
    await browser.close();
    process.exit(1);
  } catch (err) {
    clearInterval(checkInterval);
    await browser.close();
    process.exit(1);
  }
})();
