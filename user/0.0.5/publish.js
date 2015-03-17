var ghpages = require('gh-pages');
var path = require('path');

var options = {
    push: false,
    repo: 'https://' + process.env.GITHUB_TOKEN + '@github.com/getgauge/documentation.git',
    message: 'Updating docs',
    logger: function (message) {
        console.log(message);
    }
};

ghpages.publish(path.join("_book"), options, function (err) {
    console.log(err)
});