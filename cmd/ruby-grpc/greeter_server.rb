class GreeterServer < Leadster::Greeter::Service
  def create_lead(lead_req, _unused_call)
    puts "Received: #{lead_req.name}"

    Leadster::LeadReply.new(
      id: Random.rand(1..1000).to_s,
      status: "success",
      message: "Olá, #{lead_req.name}, seu email #{lead_req.email} foi cadastrado com sucesso."
      )
  end
end
