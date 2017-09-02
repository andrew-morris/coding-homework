process.stdin.setEncoding('utf8')

let lines;
let grep_string;

grep_string = process.argv[2]

process.stdin.on('readable', () => {
    const chunk = process.stdin.read();
    if (chunk !== null) {
        lines = chunk.split('\n');
        for (i in lines) {
            if (lines[i].includes(grep_string) == true) {
                process.stdout.write(`${lines[i]}\n`);
            }
        }
    }
})

