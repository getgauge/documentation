# Markdown
set :markdown_engine, :redcarpet
set :markdown,
    fenced_code_blocks: true,
    smartypants: true,
    disable_indented_code_blocks: true,
    prettify: true,
    tables: true,
    with_toc_data: true,
    no_intra_emphasis: true

# Assets
set :css_dir, 'stylesheets'
set :js_dir, 'javascripts'
set :images_dir, 'images'
set :fonts_dir, 'fonts'

# Activate the syntax highlighter
activate :syntax

activate :autoprefixer do |config|
  config.browsers = ['last 2 version', 'Firefox ESR']
  config.cascade  = false
  config.inline   = true
end

# Github pages require relative links
activate :relative_assets
set :relative_links, true

# Build Configuration
configure :build do
  # If you're having trouble with Middleman hanging, commenting
  # out the following two lines has been known to help
  activate :minify_css
  activate :minify_javascript
  # activate :relative_assets
  # activate :asset_hash
  # activate :gzip
end

helpers do
  def print_toc(sitemap)
    tree = {}
    page_title = {}

    sitemap.resources.select{|r| r.path.end_with? ".html"}.each do |r|
      current = tree
      path=r.path
      path.split("/").inject("") do |sub_path,dir|
        sub_path = File.join(sub_path, dir)
        current[sub_path] ||= {}
        current = current[sub_path]
        page_title[sub_path] = r.data.title || "TITLE NOT SET"
        sub_path
      end
    end
    p page_title
    return print_tree false, tree, page_title 
  end
  
  private
  
  def print_tree(is_subtree, node, page_title)
    c=is_subtree ? 'tocify-subheader' : 'tocify-header'
    s="<ul class='#{c}'>"
    node.each_pair do |path, subtree|
      p=path[1...path.length]
      file_name=File.basename(path)
      unless file_name=="index.html"
        if has_index(subtree)
          s="#{s}<li class='tocify-item open'><a href='#{p}'>#{page_title[path+"/index.html"]}</a></li>"
        elsif file_name.end_with?(".html")
          s="#{s}<li class='tocify-item open'><a href='#{p}'>#{page_title[path]}</a></li>"
        else
          s="#{s}<li class='tocify-item open'><a>#{file_name}</a></li>"
        end
      end
      s=s+print_tree(true, subtree, page_title) unless subtree.empty?
    end
    return "#{s}</ul>"
  end
  
  def has_index(subtree)
    subtree.keys.any? {|x| File.basename(x) == "index.html"}
  end  
end