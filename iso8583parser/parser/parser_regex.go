package parser

import "regexp"

var threadIdRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ (\d+)\|\w\|`)
var timestampRegexMatcher = regexp.MustCompile(`\d+ (\d+) \d+ \d+\|\w\|`)

/*
3006 091613057 00153921 00153928|4| -----------------------------------                                                           .
3006 091613057 00153921 00153928|4| - BIT MAP     82 38 00 00 00 00 00 00 04 00 00 00 00 00 00 00                                 .
3006 091613057 00153921 00153928|4| -----------------------------------                                                           .
3006 091613057 00153921 00153928|4| - M.T.I      : [0800]                                                                         .
3006 091613057 00153921 00153928|4| - FLD (FIELD): LENGTH : DATA                                                                  .
3006 091613057 00153921 00153928|4| -----------------------------------                                                           .
3006 091613057 00153921 00153928|4| - FLD (007)     (010)    [0630071427]                                                         .
3006 091613057 00153921 00153928|4| - FLD (011)     (006)    [821347]                                                             .
3006 091613057 00153921 00153928|4| - FLD (012)     (006)    [101427]                                                             .
3006 091613057 00153921 00153928|4| - FLD (013)     (004)    [0630]                                                               .
3006 091613057 00153921 00153928|4| - FLD (070)     (003)    [001]                                                                .
*/

var dumpPostilionBitMapRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - BIT MAP\s* (.*)\s*\.`)
var dumpPostilionMTIRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - M.T.I\s*:\s* \[?(\d*)]?\s*\.`)
var dumpPostilionFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - FLD \((\d*\.?\d*)\)\s*\(\d*\)\s*\[(.*)\s*\.`)

/*
3006 091928224 00153921 00153943|4|  FE 34 67 81 29 E1 FA 28 01 00 00 00 00 00 00 40                                              .
3006 091928224 00153921 00153943|4| ------------------------------------                                                          .
3006 091928224 00153921 00153943|4| MESSAGE_ISO_XML_FORMAT_START[1200]                                                            .
3006 091928224 00153921 00153943|4|  <?xml version="1.0"?>                                                                        .
3006 091928224 00153921 00153943|4|  <MessageInfo>                                                                                .
3006 091928224 00153921 00153943|4|  <MsgType Value="1200"/>                                                                      .
3006 091928224 00153921 00153943|4|  <Field Number="002" Value="0000000000000000"/>                                               .
3006 091928224 00153921 00153943|4|  <Field Number="003" Value="000000"/>                                                         .
3006 091928224 00153921 00153943|4|  <Field Number="004" Value="000000000000"/>                                                   .
3006 091928224 00153921 00153943|4|  <Field Number="005" Value="000000000000"/>                                                   .
3006 091928224 00153921 00153943|4|  <Field Number="006" Value="000000000000"/>                                                   .
3006 091928224 00153921 00153943|4|  <Field Number="007" Value="0000000000"/>                                                     .
3006 091928224 00153921 00153943|4|  <Field Number="011" Value="000000"/>                                                         .
3006 091928224 00153921 00153943|4|  <Field Number="012" Value="000000000000"/>                                                   .
3006 091928224 00153921 00153943|4|  <Field Number="014" Value="0000"/>                                                           .
3006 091928224 00153921 00153943|4|  <Field Number="018" Value="0000"/>                                                           .
3006 091928224 00153921 00153943|4|  <Field Number="019" Value="000"/>                                                            .
3006 091928224 00153921 00153943|4|  <Field Number="022" Value="000000000000"/>                                                   .
3006 091928224 00153921 00153943|4|  <Field Number="023" Value="000"/>                                                            .
3006 091928224 00153921 00153943|4|  <Field Number="024" Value="000"/>                                                            .
3006 091928224 00153921 00153943|4|  <Field Number="025" Value="0000"/>                                                           .
3006 091928224 00153921 00153943|4|  <Field Number="032" Value="000000"/>                                                         .
3006 091928224 00153921 00153943|4|  <Field Number="035" Value="0000000000000000000000000000000000000"/>                          .
3006 091928224 00153921 00153943|4|  <Field Number="037" Value="000000000000"/>                                                   .
3006 091928224 00153921 00153943|4|  <Field Number="040" Value="000"/>                                                            .
3006 091928224 00153921 00153943|4|  <Field Number="041" Value="00000000"/>                                                       .
3006 091928224 00153921 00153943|4|  <Field Number="042" Value="000000000000000"/>                                                .
3006 091928224 00153921 00153943|4|  <Field Number="043" Value="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"/>                       .
3006 091928224 00153921 00153943|4|  <Field Number="048" Value="fezffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff   .
3006 091928224 00153921 00153943|4|  fezfezgfregregcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccrregrg"/>     .
3006 091928224 00153921 00153943|4|  <Field Number="049" Value="000"/>                                                            .
3006 091928224 00153921 00153943|4|  <Field Number="050" Value="000"/>                                                            .
3006 091928224 00153921 00153943|4|  <Field Number="051" Value="000"/>                                                            .
3006 091928224 00153921 00153943|4|  <Field Number="052" Value="0000000000000000"/>                                               .
3006 091928224 00153921 00153943|4|  <Field Number="053" Value="00000000000000000000000000000000000000000000000      0000000000   .
3006 091928224 00153921 00153943|4|  0000000000"/>                                                                                .
3006 091928225 00153921 00153943|4|  <Field Number="055" Value="000000000000000000000000000000000000000000000000000000000000000   .
3006 091928225 00153921 00153943|4|  000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000   .
3006 091928225 00153921 00153943|4|  000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000   .
3006 091928225 00153921 00153943|4|  0000000000000"/>                                                                             .
3006 091928225 00153921 00153943|4|  <Field Number="059" Value="0000000000"/>                                                     .
3006 091928225 00153921 00153943|4|  <Field Number="061" Value="00000000000000000"/>                                              .
3006 091928225 00153921 00153943|4|  <Field Number="072" Value="00000000000000000000000000000000000000000000000000000000000000000 .
3006 091928225 00153921 00153943|4|  <Field Number="122" Value="000000000000000"/>                                                .
3006 091928225 00153921 00153943|4|  </MessageInfo>                                                                               .
*/

var dumpXmlBitMapRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s* (.*) \s*\.`)
var dumpXmlMTIRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*MESSAGE_ISO_XML_FORMAT_START\[(\d*)]\s*\.`)
var dumpXmlFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s* <Field Number="(\d{3})" Value="(.*)"/>\s*\.`)
var dumpXmlClosingMessageInfoTagRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*</MessageInfo>\s*\.`)
var dumpXmlMultiLineFieldHeaderRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*<Field Number="(\d{3})" Value="(.*)\s*\.`)
var dumpXmlMultiLineFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s* (.*)\s*\.`)
var dumpXmlMultiLineFieldClosingTagRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*(.*)"/>\s*\.`)

/*
3006 091928312 00153921 00153943|4| - BIT MAP  :  FE 36 67 81 2B E1 FA 28 01 00 00 00 04 00 00 40                                 .
3006 091928312 00153921 00153943|4| ------------------------------------                                                          .
3006 091928312 00153921 00153943|4| - M.T.I      : 1200                                                                           .
3006 091928312 00153921 00153943|4| ------------------------------------                                                          .
3006 091928312 00153921 00153943|4| - FLD (FIELD): LENGTH :  DATA                                                                 .
3006 091928312 00153921 00153943|4| ------------------------------------                                                          .
3006 091928312 00153921 00153943|4| - FLD (002) : (016) : [0000000000000000]                                                      .
3006 091928312 00153921 00153943|4| - FLD (003) : (006) : [000000]                                                                .
3006 091928312 00153921 00153943|4| - FLD (004) : (012) : [000000000000]                                                          .
3006 091928312 00153921 00153943|4| - FLD (005) : (012) : [000000000000]                                                          .
3006 091928312 00153921 00153943|4| - FLD (006) : (012) : [000000000000]                                                          .
3006 091928312 00153921 00153943|4| - FLD (007) : (010) : [0000000000]                                                            .
3006 091928312 00153921 00153943|4| - FLD (011) : (006) : [000000]                                                                .
3006 091928312 00153921 00153943|4| - FLD (012) : (012) : [000000000000]                                                          .
3006 091928312 00153921 00153943|4| - FLD (014) : (004) : [0000]                                                                  .
3006 091928312 00153921 00153943|4| - FLD (015) : (006) : [000000]                                                                .
3006 091928312 00153921 00153943|4| - FLD (018) : (004) : [0000]                                                                  .
3006 091928312 00153921 00153943|4| - FLD (019) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (022) : (012) : [000000000000]                                                          .
3006 091928312 00153921 00153943|4| - FLD (023) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (024) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (025) : (004) : [0000]                                                                  .
3006 091928312 00153921 00153943|4| - FLD (032) : (006) : [000000]                                                                .
3006 091928312 00153921 00153943|4| - FLD (035) : (037) : [0000000000000000000000000000000000000]                                 .
3006 091928312 00153921 00153943|4| - FLD (037) : (012) : [000000000000]                                                          .
3006 091928312 00153921 00153943|4| - FLD (039) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (040) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (041) : (008) : [00000000]                                                              .
3006 091928312 00153921 00153943|4| - FLD (042) : (015) : [000000000000000]                                                       .
3006 091928312 00153921 00153943|4| - FLD (043) : (040) : [0000000000000000000000000000000000000000]                              .
3006 091928312 00153921 00153943|4| - FLD (048) : (381) :                                                                         .
3006 091928312 00153921 00153943|4|        >Tag: 'P02' : ( XXXXXXXXXXXX............) : [Unkown issuer: undefined routing]         .
3006 091928312 00153921 00153943|4|        >Tag: 'P03' : ( XXXXXXXXXXXX............) : [000]                                      .
3006 091928312 00153921 00153943|4|        >Tag: 'P06' : ( XXXXXXXXXXXX............) : [000]                                      .
3006 091928312 00153921 00153943|4|        >Tag: 'P09' : ( XXXXXXXXXXXX............) : [000000]                                   .
3006 091928312 00153921 00153943|4|        >Tag: 'P10' : ( XXXXXXXXXXXX............) : [00000000000000000000000000000000000000000 .
3006 091928312 00153921 00153943|4|        >Tag: 'P13' : ( XXXXXXXXXXXX............) : [000000]                                   .
3006 091928312 00153921 00153943|4|        >Tag: 'P20' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P22' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P23' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P25' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P26' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P27' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P29' : ( XXXXXXXXXXXX............) : [00000000000000000000000000000000]         .
3006 091928312 00153921 00153943|4|        >Tag: 'P36' : ( XXXXXXXXXXXX............) : [000000]                                   .
3006 091928312 00153921 00153943|4|        >Tag: 'P64' : ( XXXXXXXXXXXX............) : [000000000]                                .
3006 091928312 00153921 00153943|4|        >Tag: 'P67' : ( XXXXXXXXXXXX............) : [000000000000000000]                       .
3006 091928312 00153921 00153943|4|        >Tag: 'P69' : ( XXXXXXXXXXXX............) : [000000000000000]                          .
3006 091928312 00153921 00153943|4|        >Tag: 'P87' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P88' : ( XXXXXXXXXXXX............) : [0]                                        .
3006 091928312 00153921 00153943|4|        >Tag: 'P95' : ( XXXXXXXXXXXX............) : [00]                                       .
3006 091928312 00153921 00153943|4|        >Tag: 'P97' : ( XXXXXXXXXXXX............) : [0000]                                     .
3006 091928312 00153921 00153943|4|        >Tag: 'O01' : ( XXXXXXXXXXXX............) : [0000]                                     .
3006 091928312 00153921 00153943|4|        >Tag: 'K03' : ( XXXXXXXXXXXX............) : [000000]                                   .
3006 091928312 00153921 00153943|4| - FLD (049) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (050) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (051) : (003) : [000]                                                                   .
3006 091928312 00153921 00153943|4| - FLD (052) : (016) : [0000000000000000]                                                      .
3006 091928312 00153921 00153943|4| - FLD (053) : (073) :                                                                         .
3006 091928312 00153921 00153943|4|        > SECURITY FORMAT          ..............: [00]                                        .
3006 091928312 00153921 00153943|4|        > PIN BLOCK FORMAT         ..............: [00]                                        .
3006 091928312 00153921 00153943|4|        > PIN ENCRYPTION KEY INDEX ..............: [000]                                       .
3006 091928312 00153921 00153943|4|        > MAC KEY INDEX            ..............: [000]                                       .
3006 091928313 00153921 00153943|4|        > SOURCE PID               ..............: [00000000]                                  .
3006 091928313 00153921 00153943|4|        > SOURCE RESOURCE          ..............: [000000]                                    .
3006 091928313 00153921 00153943|4|        > DESTINATION PID          ..............: [00000000]                                  .
3006 091928313 00153921 00153943|4|        > DESTINATION RESOURCE     ..............: [000000]                                    .
3006 091928313 00153921 00153943|4|        > USER                     ..............: [XXXXXXXXX      ]                           .
3006 091928313 00153921 00153943|4|        > PURGE TIME               ..............: [000000000]                                 .
3006 091928313 00153921 00153943|4|        > SOURCE NODE              ..............: [0000]                                      .
3006 091928313 00153921 00153943|4|        > DESTINATION NODE         ..............: [0000]                                      .
3006 091928313 00153921 00153943|4|        > PURGE TIME MS            ..............: [000]                                       .
3006 091928313 00153921 00153943|4| - FLD (055) : (132) :                                                                         .
3006 091928313 00153921 00153943|4|      > [8200]|ISO_TAG_APP_INTER_PROFILE    ....: [004] : [0000]                               .
3006 091928313 00153921 00153943|4|      > [8400]|ISO_TAG_DED_FILE_NAME        ....: [014] : [00000000000000]                     .
3006 091928313 00153921 00153943|4|      > [9500]|ISO_TAG_TVR                  ....: [010] : [0000000000]                         .
3006 091928313 00153921 00153943|4|      > [9A00]|ISO_TAG_TRANS_DATE           ....: [006] : [000000]                             .
3006 091928313 00153921 00153943|4|      > [9C00]|ISO_TAG_TRANS_TYPE           ....: [002] : [00]                                 .
3006 091928313 00153921 00153943|4|      > [5F2A]|ISO_TAG_TRANS_CUR_CODE       ....: [004] : [0000]                               .
3006 091928313 00153921 00153943|4|      > [9F02]|ISO_TAG_TRANS_AMOUNT         ....: [012] : [000000000000]                       .
3006 091928313 00153921 00153943|4|      > [9F03]|ISO_TAG_OTHER_AMOUNT         ....: [012] : [000000000000]                       .
3006 091928313 00153921 00153943|4|      > [9F09]|ISO_TAG_TERM_APP_VER_NUM     ....: [004] : [0000]                               .
3006 091928313 00153921 00153943|4|      > [9F10]|ISO_TAG_ISS_APP_DATA         ....: [014] : [00000000000000]                     .
3006 091928313 00153921 00153943|4|      > [9F1A]|ISO_TAG_TERM_COUNTRY_CODE    ....: [004] : [0000]                               .
3006 091928313 00153921 00153943|4|      > [9F1E]|ISO_TAG_IFD_SERIAL_NUM       ....: [008] : [00000000]                           .
3006 091928313 00153921 00153943|4|      > [9F26]|ISO_TAG_APP_CRYPTOGRAM       ....: [016] : [0000000000000000]                   .
3006 091928313 00153921 00153943|4|      > [9F27]|ISO_TAG_CRYPTO_INFO_DATA     ....: [002] : [00]                                 .
3006 091928313 00153921 00153943|4|      > [9F33]|ISO_TAG_TERM_CAP             ....: [006] : [000000]                             .
3006 091928313 00153921 00153943|4|      > [9F34]|ISO_TAG_CVM                  ....: [006] : [000000]                             .
3006 091928313 00153921 00153943|4|      > [9F35]|ISO_TAG_TERM_TYPE            ....: [002] : [00]                                 .
3006 091928313 00153921 00153943|4|      > [9F36]|ISO_TAG_ATC                  ....: [004] : [0000]                               .
3006 091928313 00153921 00153943|4|      > [9F37]|ISO_TAG_UNPRED_NUMBER        ....: [008] : [00000000]                           .
3006 091928313 00153921 00153943|4|      > [9F41]|ISO_TAG_TRANS_SEQ_NUM        ....: [008] : [00000000]                           .
3006 091928313 00153921 00153943|4| - FLD (059) : (010) : [0000000000]                                                            .
3006 091928313 00153921 00153943|4| - FLD (061) : (017) :                                                                         .
3006 091928313 00153921 00153943|4|        >Tag: '022' : ( N_POS_ENTRY_MODE........) : [000]                                      .
3006 091928313 00153921 00153943|4|        >Tag: '025' : ( N_POS_CONDITION_CODE....) : [00]                                       .
3006 091928313 00153921 00153943|4| - FLD (072) : (086) : [000000000000000000000000000000000000000000000000000000000000000000000] .
3006 091928313 00153921 00153943|4| - FLD (102) : (010) : [0000000000]                                                            .
3006 091928313 00153921 00153943|4| - FLD (122) : (015) : [000000000000000]                                                       .
3006 091928313 00153921 00153943|4| ------------------------------------                                                          .
*/

var dumpIsoBitMapRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - BIT MAP\s*:\s*(.*)\s*\.`)
var dumpIsoMTIRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - M.T.I\s*:\s* \[?(\d*)]?\s*\.`)
var dumpIsoFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - FLD \((\d*)\)\s*:\s*\(\d*\)\s*:\s*\[(.*)\s*\.`)
