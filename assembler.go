package main

import "fmt"
import "strings"

func findvalue(x string){
	codes:=make(map[string]int)
	codes["hlt"] = 0
	codes["ld"] = 1
	codes["sto"] = 2
	codes["ld#"] = 3
	codes["ldi"] = 4
	codes["add"] = 5
	codes["sub"] = 6
	codes["mul"] = 7
	codes["div"] = 8
	codes["jmp"] = 10
	codes["jz"] = 11
	lookup:=make(map[string]int)
	lookup["r0"] = 0
	lookup["r1"] = 1
	lookup["r2"] = 2
	lookup["r3"] = 3
	lookup["r4"] = 4
	lookup["r5"] = 5
	lookup["r6"] = 6
	lookup["r7"] = 7
	lookup["r8"] = 8
	lookup["r9"] = 9
	temp := 0
	for p, va := range codes{
		if (strings.Index(x, p) != -1){
			for kr, vr := range lookup{
				j := strings.Index(x, kr)				
				if (j != -1 ){
					balstr:=x[j+len(kr):]
					for kr2, vr2 := range lookup{
						if (strings.Index(balstr, kr2) != -1){
							printfun(p,kr,kr2,va,vr,vr2)
							temp = 1
							break
						}else if(vr2 == 6 && temp != 1){
							kr2=""
							vr2=0
							printfun(p,kr,kr2,va,vr,vr2)
						}	
					}
				}
			}
		}
	}
}

func printfun(codestr,lkstr,l2str string,codeval,lkval,l2val int){
	fmt.Println("ASSEMBLY CODE:")
	fmt.Println(codestr,lkstr,l2str)
	cval:=(codeval*10000)+(1000*lkval)+l2val
	fmt.Println("MACHINE CODE:",cval,"\n")
}

func main(){
	fmt.Println("\n\n#####ASSEMBLER#####\n")
	k:= "add r1, r2"
	findvalue(k)	
	k = "div r4,r1"
	findvalue(k)	
	k = "mul r1, r2"
	findvalue(k)	

}

/*
100 030000   go    ld#  r0,0      reg 0 will hold the sum
101 031107         ld#  r1,nums   reg 1 points to next num
102 042001   loop  ldi  r2,r1     get next number into reg 2
103 112106         jz   r2,done   if its zero we're done
104 050002         add  r0,r2     else add it to the sum
105 100102         jmp  loop      go for the next one
106 000000   done  hlt            all done. sum is in reg 0
107 000123   nums  123            the numbers to add
108 000234         234
109 000345         345
110 000000           0            zero marks the end*/
