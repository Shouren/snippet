from libnmap.process import NmapProcess
from libnmap.parser import NmapParser

nm = NmapProcess('43.241.223.0/24', '-sn -PA -PS -n')
nm.run()
report = NmapParser.parse(nm.stdout)

print report.hosts[253].is_up()
