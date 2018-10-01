@MAX_ARG_COUNT = 1

def calculate_subfactorial(input)
  case input
  when 0
    1
  when 1
    0
  else
    (input - 1) * (calculate_subfactorial(input - 1) + calculate_subfactorial(input - 2))
  end
end

def argument_count_correct?(args)
  args.length <= @MAX_ARG_COUNT
end

if not argument_count_correct?(ARGV)
  raise 'This script only takes one number!'
end

begin
  input = Integer(ARGV[0])
rescue
  raise 'This script only takes integers.'
end

print "Input: #{input.to_s}. \
Number of subfactorials: #{calculate_subfactorial(input).to_s}\n"
