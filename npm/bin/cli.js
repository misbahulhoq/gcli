#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");

const platform = process.platform;
const binName = platform === "win32" ? "git-ac.exe" : "git-ac";
const binPath = path.join(__dirname, "..", "bin", binName);

const child = spawn(binPath, process.argv.slice(2), { stdio: "inherit" });

// 4. Handle exit codes (pass success/failure back to the terminal)
child.on("close", (code) => {
  process.exit(code);
});

child.on("error", (err) => {
  console.error("❌ Failed to start git-ac:", err.message);
  console.error(`   Looking for binary at: ${binPath}`);
  process.exit(1);
});
