metadata :name => "echo",
         :description => "Choria Echo Agent",
         :author => "R.I.Pienaar <rip@devco.net>",
         :license => "Apache-2",
         :version => "1.0.0",
         :url => "https://choria.io",
         :timeout => 2

action "ping", :description => "pings a Choria server" do
    display :always

    input :message,
          :prompt  => "Message",
          :description => "Message to send",
          :type => :string,
          :validation => ".",
          :optional => true,
          :default => "ping",
          :maxlength => 128
end