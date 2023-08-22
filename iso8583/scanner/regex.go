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

const fld37DumpPostilionRegex string = `.*FLD \(037\).*\[(.*?)\]`
const fld37XmlDumpRegex string = `.*<Field Number="037" Value="(\d*)"\/>`
const fld37DumpBufferRegex string = `.*037 {\s*RRN} \d* (\d*)\s*\.`

const startDumpPostilionRegex string = `\d+ \d+ \d+ \d+\|\w\|\s*Start\s*DumpPostilion\(\)\s*\.`
const endDumpPostilionRegex string = `\d+ \d+ \d+ \d+\|\w\|\s*End\s*DumpPostilion\(\)\s*\.`

const startXmlDumpRegex string = `\d+ \d+ \d+ \d+\|\w\|\s*Start\s*DumpFile\(\)\s*\.`
const endXmlDumpRegex string = `\d+ \d+ \d+ \d+\|\w\|\s*End\s*DumpFile\(\)\s*\.`

const startDumpIso string = `\d+ \d+ \d+ \d+\|\w\|\s*Start\s*DumpIso\(\)\s*\.`
const endDumpIso string = `\d+ \d+ \d+ \d+\|\w\|\s*End\s*DumpIso\(\)\s*\.`

const startDumpTlvBuffer string = `\d+ \d+ \d+ \d+\|\w\|\s*Start\s*PrintTlvBuffer\s*\(\)\s*\.`
const endDumpTlvBuffer string = `\d+ \d+ \d+ \d+\|\w\|\s*End\s*PrintTlvBuffer\s*\(\)\s*\.`
