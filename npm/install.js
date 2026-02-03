const fs = require("fs");
const path = require("path");
const { execSync } = require("child_process");

// 1. Configuration
const packageJson = require("./package.json");
const version = packageJson.version;
const BIN_DIR = path.join(__dirname, "bin");
const BIN_NAME = "git-ac"; // The command name

// 2. Detect Platform & Arch
const platform = process.platform; // 'linux', 'darwin', 'win32'
const arch = process.arch; // 'x64', 'arm64'

console.log(`Detected platform: ${platform} (${arch})`);

// 3. Map to GitHub Release Filenames
let fileName = "";

if (platform === "win32") {
  fileName = "git-ac-windows.exe";
} else if (platform === "linux") {
  fileName = "git-ac-linux";
} else if (platform === "darwin") {
  // Mac (Intel or M1) - Assuming you use the same binary or Rosetta handles it
  fileName = "git-ac-mac";
} else {
  console.error(`❌ Unsupported platform: ${platform}`);
  process.exit(1);
}

// 4. Construct Download URL
const url = `https://github.com/misbahulhoq/git-ac/releases/download/v${version}/${fileName}`;
const destPath = path.join(
  BIN_DIR,
  platform === "win32" ? `${BIN_NAME}.exe` : BIN_NAME,
);

// 5. Download Function
async function install() {
  if (!fs.existsSync(BIN_DIR)) fs.mkdirSync(BIN_DIR);

  console.log(`⬇️  Downloading from: ${url}`);

  try {
    const response = await fetch(url);
    if (!response.ok)
      throw new Error(`Failed to fetch: ${response.statusText}`);

    const buffer = await response.arrayBuffer();
    fs.writeFileSync(destPath, Buffer.from(buffer));

    // 6. Make executable (Linux/Mac only)
    if (platform !== "win32") {
      execSync(`chmod +x "${destPath}"`);
    }

    console.log(`✅ Installed successfully to: ${destPath}`);
  } catch (err) {
    console.error("❌ Download failed:", err.message);
    process.exit(1);
  }
}

install();
