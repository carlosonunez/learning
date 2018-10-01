require 'spec_helper'

describe 'Given a method that calculates derangements' do
  context 'When we use the recursive method' do
    it 'Calculates the correct derangements' do
      @derangement_test_pairs.keys.each do |input|
        expected_value = @derangement_test_pairs[input]
        expect(Methods::Recursive.calculate_total_number_of_derangements(input)).to
          eq expected_value
      end
    end
  end
end
