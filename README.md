# DNSSHOW
Extra info on DNS commands emphasizing the education aspect

TODO:

1. Packet content breakdown


## Output example (still messy)

```
--- MESSAGE HEADER ---
                              1  1  1  1  1  1
0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      19877                    |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
| 0 |    0      | 0| 0| 1| 0| 0| 0| 0|     0    |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      1                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      0                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      0                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      0                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
--- QUESTION ---
+---------------------------------------------+
| QNAME                                       |
+---------------------------------------------|
| QTYPE                                       |
+---------------------------------------------|
| QCLASS                                      |
+---------------------------------------------+

question: &[{www.site.com. 1 1}]

question.length: 1   ns2.site.com.    142800  IN      A       x.x.x.x
Checking y.y.y.y www.site.com.

 msg is: ;; opcode: QUERY, status: NOERROR, id: 0
;; flags:; QUERY: 0, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 0

 msg is now: ;; opcode: QUERY, status: NOERROR, id: 49791
;; flags: rd; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;www.site.com.        IN       A
Populated, this is
            --------HEADER--------               
                              1  1  1  1  1  1
0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      ID                       |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|QR|   OpCode  |AA|TC|RD|RA| Z|AD|CD|   RCODE   |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                QDCOUNT/ZOCOUNT                |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                ANCOUNT/PRCOUNT                |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                NSCOUNT/UPCOUNT                |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    ARCOUNT                    |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+

 msmessage g is now: ;; opcode: QUERY, status: NOERROR, id: 49791
;; flags: rd; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;www.site.com.        IN       A

compressed: 0

ID: 49791

QR: 0

opCode: 0

Authoritative: 0

truncated: 0

recursionDesired: 1

recursionAvailable: 0

zero: 0
authenticatedData: 0

checkingDisabled: 0
rcode: 0

Question:
[{www.site.com. 1 1}]
Answer: []
Ns: []

Extra: []

 

qdcount: 1

qdcount: 0

qdcount: 0

qdcount: 0
--- MESSAGE HEADER ---
                              1  1  1  1  1  1
0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      49791                    |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
| 0 |    0      | 0| 0| 1| 0| 0| 0| 0|     0    |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      1                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      0                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      0                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      0                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
--- QUESTION ---
+---------------------------------------------+
| QNAME                                       |
+---------------------------------------------|
| QTYPE                                       |
+---------------------------------------------|
| QCLASS                                      |
+---------------------------------------------+

question: &[{www.site.com. 1 1}]

question.length: 1   www.site.com.    300     IN      A       x.x.x.x
Result: x.x.x.x

```