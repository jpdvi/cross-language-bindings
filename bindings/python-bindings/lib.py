from ctypes import *
import json

lib = cdll.LoadLibrary("bin/lib.so")
lib.DoSomething.argtypes = [c_char_p]
lib.DoSomething.restype = POINTER(c_ubyte*64)

ptr = lib.DoSomething("hello from python".encode("utf-8"))

length = int.from_bytes(ptr.contents, byteorder='little')
data = bytes(cast(ptr,
                  POINTER(c_ubyte*(64+length))).contents[64:])

jsonResponse = json.loads(data.decode())
print(jsonResponse)
