package scanner

/*
format:
date time process_id thread_id|level| message

3006 091613057 00153921 00153928|5| Read Socket OK : nRead = 51 nLenMsg=51                                                        .
3006 091613057 00153921 00153928|0| Start dump_buffer()                                                                           .
3006 091613057 00153921 00153928|D|                                      |START                                                   .
3006 091613057 00153921 00153928|D|                                      |START                                                   .
3006 091613057 00153921 00153928|D|                                      |START                                                   .
3006 091613057 00153921 00153928|D|                                      |TIMESTAMP           3006091613                          .
3006 091613057 00153921 00153928|D| 0 8 0 0 . 8 . . . . . . . . . . . .  |30 38 30 30 82 38 00 00 00 00 00 00 04 00 00 00 00 00   .
3006 091613057 00153921 00153928|D| . . 0 6 3 0 0 7 1 4 2 7 8 2 1 3 4 7  |00 00 30 36 33 30 30 37 31 34 32 37 38 32 31 33 34 37   .
3006 091613057 00153921 00153928|D| 1 0 1 4 2 7 0 6 3 0 0 0 1            |31 30 31 34 32 37 30 36 33 30 30 30 31                  .
3006 091613057 00153921 00153928|D|                                      |END                 3006091613                          .
*/

// match[1] = date
// match[2] = time
// match[3] = process_id
// match[4] = thread_id
// match[5] = level
// match[6] = message
const lineRegex string = `/(\d+) (\d+) (\d+) (\d+)\|(\w)\|(.*)\./gm`

const fld37DumpPostilionRegex string = `.*FLD \(037\).*\[(.*?)\]`
const fld37XmlDumpRegex string = `.*<Field Number="037" Value="(\d*)"\/>`
const fld37DumpBufferRegex string = `.*037 {\s*RRN} \d* (\d*)\s*\.`

// const dump_postilion_regex string = `\d+ \d+ \d+ \d+\|\w\|\s*Start\s*DumpPostilion\(\)\s*\.\n(\d+\s\d+\s\d+\s\d+\|\w\|\s*-.*\.\n)*\d+ \d+ \d+ \d+\|\w\|\s*End\s*DumpPostilion\(\)\s*\.`

// startDumpPostilionRegex the regex to match the start of dump_postilion
const startDumpPostilionRegex string = `\d+ \d+ \d+ \d+\|\w\|\s*Start\s*DumpPostilion\(\)\s*\.`

// endDumpPostilionRegex the regex to match the end of dump_postilion
const endDumpPostilionRegex string = `\d+ \d+ \d+ \d+\|\w\|\s*End\s*DumpPostilion\(\)\s*\.`

// dataDumpPostilionRegex the regex to match the data fields of dump_postilion
const dataDumpPostilionRegex string = `\d+ \d+ \d+ \d+\|\w\|\s*-+.*\.`
