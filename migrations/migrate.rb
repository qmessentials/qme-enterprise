require 'pg'
require 'fileutils'
require 'tomlrb'

config = Tomlrb.parse_file('migrations.toml')
app = ARGV[0]

server = config[app]['server']
port = config[app]['port']
dbname = config[app]['dbname']
user = ENV['DATABASE_USER'] 
password = ENV['DATABASE_PASSWORD'] 

def connect_to_db(alternate_dbname = "")
  conn = PG.connect(
    dbname: alternate_dbname || dbname,
    user: user,
    password: password,
    host: server,
    port: port
  )
  yield conn
ensure
  conn&.close
end

def create_if_not_exists
  connect_to_db('postgres') do |conn| 
    db_exists = conn.exec("select exists (select null from pg_database where datname = '#{dbname}');")
    unless db_exists 
      conn.exec "create database #{dbname};"
    end
  end
end

def get_current_schema_version
  connect_to_db do |conn|
    result = conn.exec("SELECT MAX(version_number) FROM schema_version")
    result.first['max']&.to_i || 0
  end
end

def apply_migrations(from_version)
  migrations_dir = "migrations/#{app}"
  (from_version + 1..Float::INFINITY).each do |i|
    migration_file = "#{migrations_dir}/#{i.to_s.rjust(5, '0')}.psql"
    unless File.exist?(migration_file)
      break
    end

    puts "Applying migration: #{migration_file}"
    connect_to_db do |conn|
      conn.exec(File.read(migration_file))
      conn.exec("INSERT INTO schema_version (version_number) VALUES (#{i})")
    end
  end
end

create_if_not_exists
current_version = get_current_schema_version
apply_migrations(current_version)