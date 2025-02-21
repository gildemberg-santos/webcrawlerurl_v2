this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require_relative 'main_pb'
require_relative 'main_services_pb'

def main
  name = ARGV.size > 0 ?  ARGV[0] : 'World'
  stub = Helloworld::Greeter::Stub.new('localhost:50051', :this_channel_is_insecure)
  message = stub.say_hello(Helloworld::HelloRequest.new(name: name)).message
  puts "Greeting: #{message}"
end

main
