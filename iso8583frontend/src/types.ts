export interface Message {
    fields:      { [key: string]: Field };
    mti:         Mti;
    bitmap:      string;
    raw:         string;
    logFileName: string;
    lineNumber:  string;
}

export interface Field {
    length: string;
    value:  string;
    raw:    string;
}

export interface Mti {
    version:  number;
    class:    number;
    function: number;
    origin:   number;
}

export interface IParseResult {
    id: string;
    name: string;
    date: Date;
    status: 'downloading' | 'parsing' | 'done' | 'error';
    type: 'ssh' | 'upload'
    messages: Message[];
}

export interface Search {
    mtiVersion?: string;
    mtiClass?: string;
    mtiFunction?: string;
    mtiOrigin?: string;
    bitmap?: string;
    logFiles?: string[];
    fields?: { [key: string]: string };
}

const ver1993Messages = {
    '000': 'APPROVED',
    '001': 'APPROVED WITH ID',
    '002': 'PARTIAL APPROVED',
    '003': 'APPROVED VIP',
    '004': 'APPROVED, TRACKIII UPDATE',
    '007': 'APPROVED UPDATE CHIP',
    '008': 'NO SUFFICIENT NOTES',
    '010': 'NOTES NOT AVAILABLE',
    '011': 'MAX AMOUNT 10000E',
    '045': 'FALLBACK TRN NOT ALLOWED',
    '060': 'WRONG PIN LAST RETRY',
    '061': 'SHORT NEW PIN LENGTH',
    '070': 'RECEIPT NOT AVAILABLE',
    '085': 'NO REASON TO DECLINE',
    '100': 'REJECTED',
    '101': 'EXPIRED CARD',
    '102': 'CRYPTOGRAM ERROR',
    '104': 'PRIVATE CARD',
    '106': 'PIN TRIES EXCEEDED',
    '107': 'REFER TO CARD ISSUER',
    '108': 'REFER TO ISSUER; SPECIAL COND',
    '109': 'INVALID MERCHANT',
    '110': 'INVALID AMOUNT',
    '111': 'INVALID CARD NUMBER',
    '112': 'PIN ELEMENT REQUIRED',
    '114': 'UNKNOWN CARD',
    '115': 'FUNCTION NOT AVAILABLE',
    '116': 'NO SUFFICIENT FUNDS',
    '117': 'INCORRECT PIN',
    '118': 'NO SUCH CARD',
    '119': 'TRX NOT PERMITTED',
    '120': 'TRNX NOT PERMITTED',
    '121': 'EXCEEDS WITHDRAWAL AMOUNT LIMIT',
    '122': 'SECURITY VIOLATION',
    '123': 'WITHDRAWAL FREQ. EXCEEDED',
    '125': 'CARD NOT IN SERVICE',
    '126': 'WRONG PIN FORMAT',
    '127': 'ERROR PIN LENGTH',
    '128': 'CRYPTOGRAPHIC ERROR',
    '129': 'SUSPECTED FRAUD',
    '130': 'CARD NOT YET VALID',
    '131': 'CLOSED ACCOUNT',
    '146': 'PIN TRIES ALREADY EXCEEDED',
    '150': 'INSUFFICIENT ATM BALANCE',
    '180': 'NO CREDIT ACCOUNT',
    '181': 'NO CHEQUE ACCOUNT',
    '182': 'NO SAVINGS ACCOUNT',
    '183': 'WRONG CVV',
    '184': 'INVALID DATE',
    '185': 'SERVICE CODE NOT SUPPORTED',
    '186': 'INVALID ACCOUNT(NO SUCH NUMBER)',
    '187': 'EXPIRED PIN',
    '188': 'GO MIGRATION PHASE1',
    '189': 'PIN NOT CHANGED',
    '199': 'REFUSED WITH TRACKIII UPDATING',
    '200': 'REJECTED, PICK UP CARD',
    '201': 'EXPIRED CARD, PICK UP',
    '202': 'FRAUDULENT,PICK UP',
    '204': 'PRIVATE CARD, PICK UP',
    '205': 'CALL SECURITY ACQUIRER, PICK UP',
    '206': 'PIN ENTRY TRIES EXCEEDED',
    '207': 'SPECIAL CONDITIONS, PICK UP',
    '208': 'LOST CARD, PICK UP',
    '209': 'STOLEN CARD,PICK UP',
    '210': 'SUSPECTED FRAUD,PICK UP',
    '280': 'ALTERNATE AMOUNT CANCELLED',
    '299': 'PICK UP',
    '300': 'SUCCESSFULLY PROCESSED',
    '301': 'NOT SUPPORTED BY THE ADDRESSEE',
    '302': 'RECORD NOT FOUND IN THE FILE',
    '303': 'DUPLICATE RECORD',
    '304': 'SECURE MERCH. NO AUTHENTICATION',
    '305': 'LOCKED TABLE',
    '306': 'UNSUCCESSFUL',
    '307': 'FORMAT ERROR',
    '308': 'DUPLICATE RECORD',
    '309': 'UNKNOWN FILE',
    '381': 'RECORD NOT FOUND',
    '382': 'ACTION NOT STOP. WITHDR',
    '383': 'DELETE BALANCE RECORD',
    '385': 'BALANCE INQUIRY',
    '400': 'REVOCATION ALL AUTHORIZATIONS',
    '451': 'FRAUD MONITORING REJECTION',
    '480': 'REVERSAL ACCEPTED',
    '481': 'INVALID REVERSAL',
    '482': 'ORIGINAL TRANS ALREADY REVERSED',
    '500': 'RECONCILIATION SUCCEEDED',
    '501': 'RECONCILIATION NOT ACCEPTED',
    '503': 'COUNTERS NOT AVAILABLE',
    '581': 'RECONCILIATION DONE',
    '582': 'RECONCILE. PROCESS NOT AVAILABLE',
    '600': 'INVALID NUMBER',
    '601': 'LOST/STOLEN',
    '602': 'DUPLICATION',
    '603': 'INVALID AMOUNT',
    '800': 'CUT-OFF IN PROGRESS',
    '880': 'CONNECTION NOT ACCEPTED',
    '881': 'DOWNLOAD HOT LIST',
    '882': 'UPDATE PARAMETERS FILE',
    '883': 'UPDATE APPLICATION FILE',
    '884': 'UPDATE SCREEN FILE',
    '885': 'UPDATE BIN TABLES',
    '888': 'SEND OF CUT-OF',
    '891': 'TPK TRANSFER',
    '892': 'TMK TRANSFER',
    '893': 'MAC TRANSFER',
    '898': 'ATM REBOOT',
    '899': 'OK, NO ACTION',
    '902': 'INVALID TRANSACTION',
    '903': 'START OF TRANSACTION',
    '904': 'RE ENTER TRANSACTION',
    '908': 'NO SUCH ISSUER',
    '909': 'SYSTEM FAULT',
    '911': 'TIME OUT',
    '912': 'ISSUER NOT AVAILABLE',
    '920': 'LIFE CYCLE DECLINE(MC USE ONLY)',
    '921': 'POLICY DECLINE(MC USE ONLY)',
    '922': 'FRAUD/SECURITY DECLINE(MC ONLY)',
    '992': 'NO CONNECTION',
    '993': 'PIN VERIFICATION FAULT',
    '994': 'TRANSACTION PROCESSING FAULT',
    '995': 'SERVER PROCESSING FAULT',
    '999': 'TEST',
}
const ver1987Messages = {
    '00': 'Approved or completed successfully',
    '01': 'Refer to card issuer',
    '02': 'Refer to card issuer\'s special conditions',
    '03': 'Invalid merchant',
    '04': 'Pick-up',
    '05': 'Do not honor',
    '06': 'Error',
    '07': 'Pick-up card, special condition',
    '08': 'Honour with identification',
    '09': 'Request in progress',
    '10': 'Approved for partial amount',
    '11': 'Approved (VIP)',
    '12': 'Invalid transaction',
    '13': 'Invalid amount',
    '14': 'Invalid card number (no such number)',
    '15': 'No such issuer',
    '16': 'Approved, update track 3',
    '17': 'Customer cancellation',
    '18': 'Customer dispute',
    '19': 'Re-enter transaction',
    '20': 'Invalid response',
    '21': 'No action taken',
    '22': 'Suspected malfunction',
    '23': 'Unacceptable transaction fee',
    '24': 'File update not supported by receiver',
    '25': 'Unable to locate record on file',
    '26': 'Duplicate file update record, old record replaced',
    '27': 'File update field edit error',
    '28': 'File update file locked out',
    '29': 'File update not successful, contact acquirer',
    '30': 'Format error',
    '31': 'Bank not supported by switch',
    '32': 'Completed partially',
    '33': 'Expired card',
    '34': 'Suspected fraud',
    '35': 'Card acceptor contact acquirer',
    '36': 'Restricted card',
    '37': 'Card acceptor call acquirer security',
    '38': 'Allowable PIN tries exceeded',
    '39': 'No credit account',
    '40': 'Requested function not supported',
    '41': 'Lost card',
    '42': 'No universal account',
    '43': 'Stolen card, pick-up',
    '44': 'No investment account',
    '51': 'Not sufficient funds',
    '52': 'No checking account',
    '53': 'No savings account',
    '54': 'Expired card',
    '55': 'Incorrect personal identification number',
    '56': 'No card record',
    '57': 'Transaction not permitted to cardholder',
    '58': 'Transaction not permitted to terminal',
    '59': 'Suspected fraud',
    '60': 'Card acceptor contact acquirer',
    '61': 'Exceeds withdrawal amount limit',
    '62': 'Restricted card',
    '63': 'Security violation',
    '64': 'Original amount incorrect',
    '65': 'Exceeds withdrawal frequency limit',
    '66': 'Card acceptor call acquirer\'s security department',
    '67': 'Hard capture (requires that card be picked up at ATM)',
    '68': 'Response received too late',
    '75': 'Allowable number of PIN tries exceeded',
    '90': 'Cutoff is in process (switch ending a day\'s business and starting the next. Transaction can be sent again in a few minutes)',
    '91': 'Issuer or switch is inoperative',
    '92': 'Financial institution or intermediate network facility cannot be found for routing',
    '93': 'Transaction cannot be completed. Violation of law',
    '94': 'Duplicate transmission',
    '95': 'Reconcile error',
    '96': 'System malfunction',
}

export const GetResponseMessage = (code: string): string =>  {
    return code?.length === 3 ? ver1993Messages[code as keyof typeof ver1993Messages] : ver1987Messages[code as keyof typeof ver1987Messages];
}