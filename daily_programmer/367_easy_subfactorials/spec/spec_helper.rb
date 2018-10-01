require 'subfactorials'
require 'rack/test'
require 'colorize'

module Mixins
  include Rack::Test::Methods
  def app
    SubfactorialsApp::Web::App
  end
end

RSpec.configure do |configuration|
  configuration.before(:all) do
    @derangement_test_pairs = {
      6: 265
      9: 133496
      14: 32071101049
    }
  end
end
