const walk = require('walk');
const path = require('path');
//avoid having same folder or file name
//returns a webpack entries of js files
const walkAndFindJSFiles = ()=>{
    let walker;
    const entries = {}
    const options = {
        listeners: {
          names: function (root, nodeNamesArray) {
            nodeNamesArray.sort(function (a, b) {
              if (a > b) return 1;
              if (a < b) return -1;
              return 0;
            });
          },
          file: function (root, fileStats, next) {
            const fileExtension = path.extname(fileStats.name)
            const name = path.basename(fileStats.name, fileExtension)
          
            if(fileExtension != ".js") next()
               entries[name] = root + "/" + fileStats.name
            
            next();
          },
          errors: function (root, nodeStatsArray, next) {
            next();
          },
        },
        filters:['dist']
      };
      
      walker = walk.walkSync(path.resolve(__dirname,'web/static/js/'), options);
      return entries
}
const entries = walkAndFindJSFiles()
module.exports = entries