this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require_relative 'main_pb'
require_relative 'main_services_pb'

def main
  name = ARGV.size > 0 ?  ARGV[0] : ""
  email = ARGV.size > 1 ?  ARGV[1] : ""
  phone = ARGV.size > 2 ?  ARGV[2] : ""

  stub = Leadster::Greeter::Stub.new('localhost:50051', :this_channel_is_insecure)
  response = stub.create_lead(Leadster::LeadRequest.new(name: name, email: email, phone: phone))
  id = response.id
  status = response.status
  message = response.message

  puts "ID: #{id}"
  puts "Status: #{status}"
  puts "Message: #{message}"
end

main
