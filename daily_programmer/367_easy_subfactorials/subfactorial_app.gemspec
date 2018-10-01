gem_author = 'Carlos Nunez'
gem_email = 'dev@carlosnunez.me'

Gem::Specification.new do |spec|
  spec.name           = 'subfactorial-app'
  spec.version        = File.read('VERSION')
  spec.authors        = [gem_author]
  spec.licenses       = ['MIT']
  spec.email          = [gem_email]
  spec.summary        = 'A way too enterprise-y version of the subfactorial Reddit problem.'
  spec.description    = 'A way too enterprise-y version of the subfactorial Reddit problem.'
  spec.files          = Dir['{bin,lib}/**/*', 'LICENSE', 'README.md']
  spec.executables    = ['subfactorial-app']
  spec.require_paths  = ['lib']
end
