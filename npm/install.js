const binwrap = require("binwrap");
const packageInfo = require("./package.json");
const version = packageInfo.version;
const root = `https://github.com/misbahulhoq/git-ac/releases/download/v${version}`;

module.exports = binwrap({
  dirname: __dirname,
  binaries: ["git-ac"],
  urls: {
    // Windows 64-bit
    "win32-x64": root + "/git-ac-windows.exe",

    // Linux 64-bit
    "linux-x64": root + "/git-ac-linux",

    // Mac (Intel)
    "darwin-x64": root + "/git-ac-mac",

    // Mac (M1/M2/Apple Silicon)
    "darwin-arm64": root + "/git-ac-mac",
  },
});
