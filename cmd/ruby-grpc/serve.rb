this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require_relative 'main_pb'
require_relative 'main_services_pb'
require_relative 'greeter_server'

def main
  s = GRPC::RpcServer.new
  s.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
  s.handle(GreeterServer)
  s.run_till_terminated_or_interrupted(['int', 'SIGTERM'])
end

main
