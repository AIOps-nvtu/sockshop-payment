from util.Docker import Docker
from util.Api import Api
import os
import unittest

class Dredd:
    image = 'weaveworksdemos/openapi'
    container_name = ''
    def test_against_endpoint(self, service, api_endpoint, links=[], env=[], dump_streams=False):
        self.container_name = Docker().random_container_name('openapi')
        command = ['docker', 'run',
                   '-h', 'openapi',
                   '--name', self.container_name,
                   '-v', "{0}:{1}".format(os.getcwd() + "/api-spec/", "/tmp/specs/")]
        
        if links != []:
            [command.extend(["--link", x]) for x in links]
            
        if env != []:
            [command.extend(["--env", "{}={}".format(x[0], x[1])]) for x in env]
            
        command.extend([Dredd.image,
                        "/tmp/specs/{0}.json".format(service),
                        api_endpoint,
                        "-f",
                        "/tmp/specs/hooks.js".format(service)])
        out = Docker().execute(command, dump_streams=dump_streams)
 
        Docker().kill_and_remove(self.container_name)
        return out
