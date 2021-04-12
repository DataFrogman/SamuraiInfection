## SamuraiInfection
Hey Samurai, some corpo went catatonic in a local braindance den.  His implants were all nonfunctional and it looks like something wiped him, unfortunately he was running some sort of archaic PortholeOS so we can't do anything to recover him.  The halo however survived, looks like it was write protected before use.  We pulled this off of it, see what you can make of it.

Author: DataFrogman

# Instructions for CA

The above text and the infection file are to be provided, the rest is from creating the challenge

# Solution

When you open the provided infection file you are confronted with a wall of text.

> CATC CCGG CA A TTTT TTTT A A CAAA A A A A A A A A A A A GAAA A GTGG TG GAC GTGA CCCA CGGA CTAG CGTT GAA CGAT CTCA GAA CTCC CGTG CACA CATT CGCA CGCC GCA A A A CGCA GACG A A A A GAG G A TGAA A A TAA CCTG CCC A A A A A CG A A A A A CTC A A A CGAA GAAC...

Looking at this, the first thought is intended to be DNA.  There are the four bases: adenine, cytosine, guanine, and thymine.  However, this is just a distraction, it is actually base four encoding.  With the groupings it is a safe guess that A == 0.
The groupings are frequently less than 4 chars per nibble, so there must be some sort of compression in place.  Let's check if it was just a simple conversion and 0 stripping, first convert each base4 nibble into base2, once we have the base2 we can pad it out with leading 0s (that way we don't change the value of the nibble), piece it together to see what we have.

```
  /mnt/f/L/Documents/SamuraiInfection/SamuraiInfection   master ❯ hexdump SamuraiInfection.exe -C | head
00000000  4d 5a 90 00 03 00 04 00  00 00 00 00 ff ff 00 00  |MZ..............|
00000010  8b 00 00 00 00 00 00 00  40 00 00 00 00 00 00 00  |........@.......|
00000020  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000030  00 00 00 00 00 00 00 00  00 00 00 00 80 00 00 00  |................|
00000040  0e 1f ba 0e 00 b4 09 cd  21 b8 01 4c cd 21 54 68  |........!..L.!Th|
00000050  69 73 20 70 72 6f 67 72  61 6d 20 63 61 6e 6e 6f  |is program canno|
00000060  74 20 62 65 20 72 75 6e  20 69 6e 20 44 4f 53 20  |t be run in DOS |
00000070  6d 6f 64 65 2e 0d 0d 0a  24 00 00 00 00 00 00 00  |mode....$.......|
00000080  50 45 00 00 64 86 03 00  00 00 00 00 00 74 19 00  |PE..d........t..|
00000090  00 00 00 00 f0 00 22 02  0b 02 03 00 00 e0 07 00  |......".........|
```
Now we have an executable!!

With an executable the first step is to check for packing, unfortunately tools return an unknown packer.  If you have experience in this area you might notice that the UPX packer flags are mangled, fixing them and unpacking gives you a simple Go binary, otherwise you have to reverse the binary the long way.

Performing analysis on the binary will show two high entropy strings and some XOR operations.  

>82, 51, 100, 95, 72, 51, 114, 114, 49, 110, 103

>0, 122, 48, 12, 13, 112, 9, 59, 92, 30, 11, 18, 93, 16, 106, 23, 115, 0, 65, 110, 8, 18, 60, 108, 39, 23, 120, 3, 63, 15

One of these strings is `R3d_H3rr1ng` so that isn't it.  Running a simple XOR operation does gives us `RITSEC{Impl@nt5_@r3_fun_CH00M}`.  The XOR is technically null preserving but that isn't visible when you are just dealing with the flag, if you try to XOR against encrypted files generated by the program you won't get the flag without performing a null preserving XOR operation.

Flag: RITSEC{Impl@nt5_@r3_fun_CH00M}