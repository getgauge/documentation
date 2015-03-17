require "git"

g = Git.open(Dir.pwd, :log => Logger.new(STDOUT))