##GOLANG Assembler
- Converting assembly code to machine code

##Assembly Codes

- 100 030000   go    ld#  r0,0      reg 0 will hold the sum
- 101 031107         ld#  r1,nums   reg 1 points to next num
- 102 042001   loop  ldi  r2,r1     get next number into reg 2
- 103 112106         jz   r2,done   if its zero we're done
- 104 050002         add  r0,r2     else add it to the sum
- 105 100102         jmp  loop      go for the next one
- 106 000000   done  hlt            all done. sum is in reg 0
- 107 000123   nums  123            the numbers to add
- 108 000234         234
- 109 000345         345
- 110 000000           0            zero marks the end

##Python code link

[click here](http://www.openbookproject.net/courses/python4fun/assembler.html)
