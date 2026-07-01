const fs = require('fs');
const path = require('path');

function processFile(filePath) {
  let content = fs.readFileSync(filePath, 'utf8');

  // Remove HTML comments <!-- ... -->
  content = content.replace(/<!--[\s\S]*?-->/g, '');

  // Remove JS single line comments that start with // (but avoiding http:// or https://)
  // This regex matches // followed by anything until end of line, but only if not preceded by :
  content = content.replace(/(?<!:)\/\/.*$/gm, '');

  // Trim empty lines that might have been left behind
  content = content.replace(/^\s*[\r\n]/gm, '');

  fs.writeFileSync(filePath, content, 'utf8');
  console.log('Processed:', filePath);
}

function walkDir(dir) {
  const files = fs.readdirSync(dir);
  for (const file of files) {
    const fullPath = path.join(dir, file);
    if (fs.statSync(fullPath).isDirectory()) {
      walkDir(fullPath);
    } else if (fullPath.endsWith('.vue') || fullPath.endsWith('.js')) {
      processFile(fullPath);
    }
  }
}

walkDir(path.join(__dirname, '../../frontend/src'));
