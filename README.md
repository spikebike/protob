Small example and test case for github.com/spikebike/backups-done-right.  Currently uses TLS and protobufs to allow client and server to communicate.

Todo:

* Consider https://github.com/capnproto/
* Send list of []byte for sha256 checksum of blob and []byte for blob contents instead of two int32s.
* Tweak client to record and print the server's TLS public key
* Tweak server to record and print the client's TLS public key
  
