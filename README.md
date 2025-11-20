# devfest2025

Example code for DevFest 2025 - Faro

This repository aims at explaining the basic concepts of blockchain
development. It is by no means comprehensive or cryptographically correct. It
is a working example that spiritually matches Bitcoin design.

It is broken up in 4 lessons and uses common code found in the
`bitcoin/bitcoin.go` directory

This code is meant to be small, readable and used for educational purposes
only. It is not cryptographically sound. It is truly only meant as an example
only.

## Lesson 1

This lesson goes over the basic theory of creating a blockchain.

```
go run lesson1/lesson1.go
```

This code pauses and expects the user to hit enter to continue to the next
part.

First we create a blockchain with 5 blocks.
```
inserted block 0 - 5b6fb58e61fa475939767d68a446f97f1bff02c0e5935a3ea8bb51e6515783d8
inserted block 1 - 3d90e6ebf192a3b069d48d30517feafb7805a6ad2d4eff181cc870f15bd3587f
inserted block 2 - 745f7101012bf26ffdc9e0418a3aeec73fff4ba01bc4f4765c210c2b0e36af06
inserted block 3 - f9a267a1e76e0dd52c80124397b8aa9c4c4b80718014f0c88dc1daae868b68ec
inserted block 4 - 3bb92e171aeb6fe0d9d817154c04cb3afc4ab85fc90b16bde800d5dc9b637228
```

Print blockchain in reverse order.
```
block 4 - 3bb92e171aeb6fe0d9d817154c04cb3afc4ab85fc90b16bde800d5dc9b637228
block 3 - f9a267a1e76e0dd52c80124397b8aa9c4c4b80718014f0c88dc1daae868b68ec
block 2 - 745f7101012bf26ffdc9e0418a3aeec73fff4ba01bc4f4765c210c2b0e36af06
block 1 - 3d90e6ebf192a3b069d48d30517feafb7805a6ad2d4eff181cc870f15bd3587f
block 0 - 5b6fb58e61fa475939767d68a446f97f1bff02c0e5935a3ea8bb51e6515783d8
```

Create 2 forked blocks.
```
inserted block 4' - ce36fac030050922365572a2c8c5e84191635dc8637577172e4c2c55dff1ad73
inserted block 5' - a073d59e64e8c834bd1f89e46ceaa3877ce198063474eb0973d313aea9b29246
total blockchain length 7
```

Print forked chain.
```
block 5 - a073d59e64e8c834bd1f89e46ceaa3877ce198063474eb0973d313aea9b29246
block 4 - ce36fac030050922365572a2c8c5e84191635dc8637577172e4c2c55dff1ad73
block 3 - f9a267a1e76e0dd52c80124397b8aa9c4c4b80718014f0c88dc1daae868b68ec
block 2 - 745f7101012bf26ffdc9e0418a3aeec73fff4ba01bc4f4765c210c2b0e36af06
block 1 - 3d90e6ebf192a3b069d48d30517feafb7805a6ad2d4eff181cc870f15bd3587f
block 0 - 5b6fb58e61fa475939767d68a446f97f1bff02c0e5935a3ea8bb51e6515783d8
```

## Lesson 2

This lesson goes over transactions creation and the importance of encoding.

```
go run lesson2/lesson2.go
```

Create transaction and encode it.
```
txid   : b369bc12d68657aebb7fd710fd10075494ad797f0abab9f529b4becc633ddcac
payload:
00000000  61 6c 69 63 65 00 00 00  00 00 00 00 00 00 00 00  |alice...........|
00000010  00 00 00 00 62 6f 62 00  00 00 00 00 00 00 00 00  |....bob.........|
00000020  00 00 00 00 00 00 00 00  00 00 00 32              |...........2|
```

Note that the encoding is required in order to be able to generate a hash of
the encoding. This is a very important concept in blockchain design.

## Lesson 3

This lesson provides a basic Proof-of-Work example. This is also known as
"mining".

```
go run lesson3/lesson3.go
```

Mining example.
```
mining   : hashes 27521832 time 1.645397687s hashes/second 16726553.232355
nonce    : 0x1a3f328
merkle   : abcc34dcb479e7bc1da80351a04befc933c41e0eb8fda88fa0ece6095edea8c2
blockhash: 000000719af43f7b87696dc8241d1485a290ffe451c7bd942b61e61d043c469f
block    :
00000000  00 00 00 01 67 65 6e 65  73 69 73 68 61 73 68 25  |....genesishash%|
00000010  25 25 25 25 25 25 25 25  25 25 25 25 25 25 25 25  |%%%%%%%%%%%%%%%%|
00000020  25 25 25 25 ab cc 34 dc  b4 79 e7 bc 1d a8 03 51  |%%%%..4..y.....Q|
00000030  a0 4b ef c9 33 c4 1e 0e  b8 fd a8 8f a0 ec e6 09  |.K..3...........|
00000040  5e de a8 c2 69 1f 01 40  00 00 00 18 01 a3 f3 28  |^...i..@.......(|
00000050  6d 69 6e 65 72 00 00 00  00 00 00 00 00 00 00 00  |miner...........|
00000060  00 00 00 00 61 6c 69 63  65 00 00 00 00 00 00 00  |....alice.......|
00000070  00 00 00 00 00 00 00 00  00 00 00 32              |...........2|
```

This shows mining statistics and the encoded block. Note that at the end of the
block you see the transaction appended. Also note the number of zeroes that
prefix the block hash.

## Lesson 4

This lesson puts all other lessons together in a practical example. It shows
essentially how bitcoin works albeit with gross simplification.

You can run all examples using the `go run` command.

```
go run lesson4/lesson4.go
```

Alice mines several blocks as shown by the updated balances. As the blockchain
progresses Alice starts sending 20 Satoshis to Bob.
```
mined block 0 - 0000004078ad5951cf6ac59be6ce7ec34ee35299dcf862fae3b270c2f574b79a alice balance 50
mined block 1 - 00000068e2c2549bfa4c7c496c457b1f1c4d3519832774a9db70f40c44f2ba45 alice balance 100
mined block 2 - 000000ba84b4cf93d9c1d57b85fd695d4bb251f35b07cac37d45e4804bdc6ff0 alice balance 150
mined block 3 - 00000072b8bcf17273281ae71c542c84249ebe277562d839277559c587e9ca74 alice balance 200
mined block 4 - 0000005b73474b8f688e9db39cdf819df80420910655a5faa96e84465222386f alice balance 250
mined block 5 - 0000003b7fa87d4f9696499505d83c3477d3fb7bee4ace2076fc7de0d53d3b92 alice balance 280 bob balance 20
mined block 6 - 000000ce19565790511efcd25b489122622fc061d36c82ca48ff29fd8a954a6a alice balance 310 bob balance 40
mined block 7 - 0000006084af6f9c6a06a4f09274a4fa34d6a366c9d315b5be339854b5cb11a2 alice balance 340 bob balance 60
mined block 8 - 000000cb6f6e86e0b0b8f320111cde75ad4d637c9f26b704b8a7e0496bf36a1e alice balance 370 bob balance 80
mined block 9 - 0000004038f9cd37e745a7e14ca70c100b792c00db5010bf12cfda53735fd9c7 alice balance 400 bob balance 100
```

This example, in spirit, shows how bitcoin progresses through the state machine
all the way through broadcasting a block.
