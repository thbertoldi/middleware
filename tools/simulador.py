import asyncio
from optparse import OptionParser

parser = OptionParser()
parser.add_option("--mf", action="store_true", dest="mono")
parser.add_option("--tf", action="store_true", dest="tri")
parser.add_option("--cm", action="store_true", dest="carbon")

(options, args) = parser.parse_args()

async def runner(model):
    await asyncio.sleep(5)

if options.mono:
    print("Simulando Dispositivo MF")
if options.tri:
    print("Simulando Dispositivo TF")
if options.carbon:
    print("Simulando Dispositivo CM")
