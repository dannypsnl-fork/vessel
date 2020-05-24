Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"

  config.vm.synced_folder "./", "/home/vagrant/vessel"

  config.vm.provider "virtualbox" do |vb|
      vb.memory = "4096"
  end
end
