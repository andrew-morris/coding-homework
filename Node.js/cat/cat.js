process.stdin.setEncoding('utf8')

var fs = require('fs')
let file_to_read;

file_to_read = process.argv[2]

fs.readFile(file_to_read, 'utf8', function(err, contents) {
    process.stdout.write(contents);
});
