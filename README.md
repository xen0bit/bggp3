# bggp3

```
MediaInfo Command line, 
MediaInfoLib - v21.09
```

`SIGSEGV.PC.7fd6f75f90b6.STACK.c2262cd14.CODE.1.ADDR.55a6b1095000.INSTR.cmpb___$0x3c,(%rdx).fuzz`

```hexdump
00000000: 6761 7262 6167 650a fffb 50c4 0000 0000  garbage...P.....
00000010: 0000 0000 0000 0000 0000 0000 0049 6e66  .............Inf
00000020: 6f00 0000 0f00 0000 0b00 0009 ca00 1717  o...............
00000030: 1717 1717 1717 172e 2e2e 2e2e 2e2e 2e2e  ................
00000040: 4545 4545 4545 4545 455d 5d5d 5d5d 5d5d  EEEEEEEEE]]]]]]]
00000050: 5d5d 7474 7474 7474 7474 748b 8b8b 8b8b  ]]ttttttttt.....
00000060: 8b8b 8b8b a2a2 a2a2 a2a2 a2a2 a2ba baba  ................
00000070: baba baba baba d1d1 d1d1 d1d1 d1d1 d1e8  ................
00000080: e8e8 e8e8 e8e8 e8e8 ffff ffff ffff ffff  ................
00000090: ff00 0000 394c 414d 4533 2e39 3872 01a5  ....9LAME3.98r..
000000a0: 0000 0000 2e1d 0000 1440 2404 4942 0000  .........@$.IB..
000000b0: 4000 0009 cae1 88d8 c400 0000 0000 0000  @...............
000000c0: 00e8 e8e8 e8e8 ffff ffff ffff ffff ff00  ................
000000d0: 0000 394c 414d 4533 2e39 3872 01a5 0000  ..9LAME3.98r....
000000e0: 0000 2e1d 0000 1440 2404 4942 0000 4000  .......@$.IB..@.
000000f0: 0009 cae1 88d8 c400 0000 0000 0000 0000  ................
00000100: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000110: 0000 0000 0000 fffb 50c4 0000 2706 6423  ........P...'.d#
00000120: 5c74 9100 01fa 21f0 7739 5203 0000 cb96  \t....!.w9R.....
00000130: 68c8 c562 b159 3c88 0018 261b 2746 8d1a  h..b.Y<...&.'F..
00000140: 31e1 ff13 83e0 f838 081c fc1c 040e 7f94  1......8........
00000150: 043f aac0 7ffe f139 f2ef ffff 28ee 0860  .?.....9....(..`
00000160: 381c 0e07 0381 c0e0 7028 1000 000a 2815  8.......p(....(.
00000170: 0573 f992 3f7d 1ae4 ed33 c98c cc34 c933  .s..?}...3...4.3
00000180: 2a53 ef6b 6197 80c2 0621 00d0 1006 0340  *S.ka....!.....@
00000190: d5c6 2d3d 8422 0121 a059 d034 36f8 0403  ..-=.".!.Y.46...
000001a0: c3f7 04c5 6062 4158 9c3f f005 1680 c050  ....`bAX.?.....P
000001b0: 8081 20e0 a980 2863 ffc1 11b4 72c0 8858  .. ...(c....r..X
000001c0: 3690 c5c1 6162 fbff f166 0c52 9108 9178  6...ab...f.R...x
000001d0: 6f9d 179f fff9 26b2 e2cf 1f93 af4a 8005  o.....&......J..
000001e0: 2b1a c922 6dba 0cff fb52 c404 800b 8109  +.."m....R......
000001f0: 8bbc b500 3179 28ae 3cd2 8ee5 9dff e15b  ....1y(.<......[
00000200: 3c                                       <
```

Does not crash: MediaInfoLib - v22.06

```
cd /
apt install -y wget p7zip-full git automake autoconf libtool pkg-config make g++ zlib1g-dev
wget https://old.mediaarea.net/download/source/mediainfo/21.09/mediainfo_21.09_AllInclusive.7z
7za x mediainfo_21.09_AllInclusive.7z
cd /mediainfo_AllInclusive/ZenLib/Project/GNU/Library
./autogen.sh
./configure --enable-debug
make
make install
cd /mediainfo_AllInclusive/MediaInfoLib/Project/GNU/Library
./autogen.sh
./configure --enable-debug
make
make install
cd /mediainfo_AllInclusive/MediaInfo/Project/GNU/CLI
./autogen.sh
./configure --enable-debug
make
make install
cp /usr/local/lib/libmediainfo.so.0 /usr/lib/
cp /usr/local/lib/libzen.so.0 /usr/lib/
mediainfo
```

`gdb --args mediainfo crash.mp3`

starti
continue
where
list

```
tfsxml_next (priv=0x55d34fb28de0, b=0x7ffe3a017ea0) at ../../../Source/ThirdParty/tfsxml/tfsxml.c:76
76              switch (*priv->buf)
(gdb) where
#0  tfsxml_next (priv=0x55d34fb28de0, b=0x7ffe3a017ea0) at ../../../Source/ThirdParty/tfsxml/tfsxml.c:76
#1  0x00007fd0b0aaf422 in MediaInfoLib::file_adm_private::parse (this=0x55d34fb28de0)
    at ../../../Source/MediaInfo/Audio/File_Adm.cpp:466
#2  0x00007fd0b0ab4bf7 in MediaInfoLib::File_Adm::FileHeader_Begin (this=0x55d34fb4cef0)
    at ../../../Source/MediaInfo/Audio/File_Adm.cpp:851
#3  0x00007fd0b095d3d0 in MediaInfoLib::File__Analyze::FileHeader_Manage (this=0x55d34fb4cef0)
    at ../../../Source/MediaInfo/File__Analyze.cpp:2134
#4  0x00007fd0b095982f in MediaInfoLib::File__Analyze::Open_Buffer_Continue_Loop (this=0x55d34fb4cef0)
    at ../../../Source/MediaInfo/File__Analyze.cpp:1082
#5  0x00007fd0b0956f13 in MediaInfoLib::File__Analyze::Open_Buffer_Continue (this=0x55d34fb4cef0, 
    ToAdd=0x55d34fb6da80 "garbage\n\377\373P\304", ToAdd_Size=513)
    at ../../../Source/MediaInfo/File__Analyze.cpp:720
#6  0x00007fd0b0a34918 in MediaInfoLib::MediaInfo_Internal::Open_Buffer_Continue (this=0x55d34fb407e0, 
    ToAdd=0x55d34fb6da80 "garbage\n\377\373P\304", ToAdd_Size=513)
    at ../../../Source/MediaInfo/MediaInfo_Internal.cpp:1655
#7  0x00007fd0b0ead8d9 in MediaInfoLib::Reader_File::Format_Test_PerParser_Continue (this=0x55d34faf2de0, 
    MI=0x55d34fb407e0) at ../../../Source/MediaInfo/Reader/Reader_File.cpp:766
#8  0x00007fd0b0eac001 in MediaInfoLib::Reader_File::Format_Test_PerParser (this=0x55d34faf2de0, 
    MI=0x55d34fb407e0, File_Name=L"crash.mp3") at ../../../Source/MediaInfo/Reader/Reader_File.cpp:313
#9  0x00007fd0b0a1660c in MediaInfoLib::MediaInfo_Internal::ListFormats (this=0x55d34fb407e0, 
    File_Name=L"crash.mp3") at ../../../Source/MediaInfo/MediaInfo_File.cpp:1176
#10 0x00007fd0b0eab743 in MediaInfoLib::Reader_File::Format_Test (this=0x55d34faf2de0, MI=0x55d34fb407e0, 
    File_Name=L"crash.mp3") at ../../../Source/MediaInfo/Reader/Reader_File.cpp:230
#11 0x00007fd0b0a31c3b in MediaInfoLib::MediaInfo_Internal::Entry (this=0x55d34fb407e0)
    at ../../../Source/MediaInfo/MediaInfo_Internal.cpp:1352
#12 0x00007fd0b0a2de4a in MediaInfoLib::MediaInfo_Internal::Open (this=0x55d34fb407e0, 
    File_Name_=L"crash.mp3") at ../../../Source/MediaInfo/MediaInfo_Internal.cpp:1108
#13 0x00007fd0b0a439b1 in MediaInfoLib::MediaInfoList_Internal::Entry (this=0x55d34fb3f260)
    at ../../../Source/MediaInfo/MediaInfoList_Internal.cpp:212
--Type <RET> for more, q to quit, c to continue without paging--
#14 0x00007fd0b0a43335 in MediaInfoLib::MediaInfoList_Internal::Open (this=0x55d34fb3f260, 
    File_Name=L"crash.mp3", Options=MediaInfoLib::FileOption_Nothing)
    at ../../../Source/MediaInfo/MediaInfoList_Internal.cpp:148
#15 0x00007fd0b0a42538 in MediaInfoLib::MediaInfoList::Open (this=0x55d34faf33e0, File=L"crash.mp3", 
    Options=MediaInfoLib::FileOption_Nothing) at ../../../Source/MediaInfo/MediaInfoList.cpp:118
#16 0x000055d34ef430a8 in Core::Menu_File_Open_Files_Continue (this=0x7ffe3a0197e0, FileName=L"crash.mp3")
    at ../../../Source/Common/Core.cpp:172
#17 0x000055d34ef3f211 in main (argc=2, argv_ansi=0x7ffe3a019978) at ../../../Source/CLI/CLI_Main.cpp:154
```

```
71      int tfsxml_next(tfsxml_string* priv, tfsxml_string* b)
72      {
73          int level = priv->flags ? 1 : 0;
74          while (priv->len)
75          {
76              switch (*priv->buf)
77              {
78              case '<':
79                  next_char(priv);
80                  if (priv->len && *priv->buf == '?')
```

```
gdb) print priv
$1 = (tfsxml_string *) 0x55d8da342de0
(gdb) pt
type = struct tfsxml_string {
    char *buf;
    int len;
    int flags;
} *
```

https://github.com/MediaArea/MediaInfoLib/commit/2ea23833d3eb0333c73fdf3e372c27a3384b342c

----

https://wimlib.net/downloads/wimlib-1.13.5.tar.gz

```
apt install wget libxml2-dev ntfs-3g-dev libfuse-dev libattr1-dev
```

```
./configure CC="hfuzz-gcc" CXX="hfuzz-g++" --disable-shared
```

```
honggfuzz -i in -z -- /usr/local/bin/wimlib-imagex verify ___FILE___
```