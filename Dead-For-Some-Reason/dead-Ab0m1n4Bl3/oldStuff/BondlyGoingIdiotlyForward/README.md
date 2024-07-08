# README.md

On the figuring out what i needed for the Relevant Box and just exploring away

#### GOALS

- Go full Mechanicus ASM your GASM
- Avoid C++ to prevent try to just be asm only just use everyones pointers 
- Learn more C -> 
- Learn more pwntools -> use it to make aid my binary patching the nice binary with my bad asm 
- Learn more asm -> binary patch my bad asm into 
- ... gurl and the polymorpheus dream of golang strings being the bestest fastest and reliable strings 


Inspired by https://cocomelonc.github.io/tutorial/2021/10/27/windows-shellcoding-1.html to write asm that was stack based which is easy on old machines so not viable but that 

Trying to be somewhere closer to the blonde doom guy and forgot my ctrl+s because why no :wq brain
but essentially it was the same as link except

we would combine with the asm from pwntools instead of have a seperate binarey to find these if this were window 7...the asnwer is https://idafchev.github.io/exploit/2017/09/26/writing_windows_shellcode.html, but I 



## pwntools utility to get into pwntools and the asm of doom
Make a more 

- Add conditionality to make pack all the things 
- With character cleverness
- slice a file by ws,tabs, newlines 
- split on 

```
from pwn import *
# inspired by https://cocomelonc.github.io/tutorial/2021/10/27/windows-shellcoding-1.html
# slapping phind till it wrote this for me

def reverse_and_hexify(input):
  chunks = [input[i:i+4][::-1] for i in range(0, len(input), 4)]
  return chunks[::-1]

input = sys.argv[1]
print(' '.join([p32(chunk)[::-1] for chunk in reverse_and_hexify(input)]))


```

Use these for finding peb kernel32.dll and etc...

https://docs.pwntools.com/en/stable/shellcraft/amd64.html#module-pwnlib.shellcraft.amd64.windows
https://github.com/Gallopsled/pwntools/blob/2e09b7dd91/pwnlib/shellcraft/templates/amd64/windows/kernel32base.asm


The answers are in https://idafchev.github.io/exploit/2017/09/26/writing_windows_shellcode.html


## Barbarian Ummencryption Maybes - see the `encrypt.c` 

AES - https://en.wikipedia.org/wiki/Advanced_Encryption_Standard

KeyExpansion – round keys are derived from the cipher key using the AES key schedule. AES requires a separate 128-bit round key block for each round plus one more.
Initial round key addition:

    AddRoundKey – each byte of the state is combined with a byte of the round key using bitwise xor.

9, 11 or 13 rounds:

    SubBytes – a non-linear substitution step where each byte is replaced with another according to a lookup table.
    ShiftRows – a transposition step where the last three rows of the state are shifted cyclically a certain number of steps.
    MixColumns – a linear mixing operation which operates on the columns of the state, combining the four bytes in each column.
    AddRoundKey

Final round (making 10, 12 or 14 rounds in total):

    SubBytes
    ShiftRows
    AddRoundKey




Phind - How to do bitmap in memory with 8x8 bits in c and do stuff with it
```c
#include <stdio.h>
#include <stdlib.h>

#define SIZE 8

void set_bit(unsigned char *bitmap, int pos, int value) {
   if (value == 1) {
       bitmap[pos / 8] |= (1 << (pos % 8));
   } else {
       bitmap[pos / 8] &= ~(1 << (pos % 8));
   }
}

int get_bit(const unsigned char *bitmap, int pos) {
   return (bitmap[pos / 8] >> (pos % 8)) & 1;
}

int main() {
   unsigned char *bitmap = malloc(SIZE * sizeof(unsigned char));

   // Set the 3rd bit to 1
   set_bit(bitmap, 2, 1);

   // Get the 3rd bit
   int bit = get_bit(bitmap, 2);
   printf("%d\n", bit); // Outputs: 1

   free(bitmap);

   return 0;
}


```


