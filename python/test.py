from ctypes import *

Promet = CDLL("promet.so")

whats_my_name = Promet.WhatsMyName
whats_my_name.argtypes = None
whats_my_name.restype = c_char_p

go_env = Promet.GoEnv
go_env.restype = c_char_p

events = Promet.Events
events.argtypes = None
events.restype = [c_char_p, c_char_p]


print "Hello", whats_my_name()
print "Env", go_env()
print "Events", events()
