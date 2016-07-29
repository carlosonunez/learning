private_key 'vagrant_host' do
  path default['setup']['private_key_dir']
  public_key_path default['setup']['private_key_dir']
  type 'rsa'
  action :create
end
