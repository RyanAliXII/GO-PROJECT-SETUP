// const walk = require('walk');
const fs = require('fs')

// const getEntries = (filesLocation)=>{
//     const entries = {}
//     const relativePathFilesLoc = filesLocation.replace("./", "")
//     const options = {
//         listeners: {
//           file: function (root, fileStats, next) {
//             folder = root.replace(`${filesLocation}`,"")
//             const entryKey = `${folder}/${fileStats.name}`
//             entries[entryKey] = `./${relativePathFilesLoc}${entryKey}`
//           },
       
//         },
//       };
    
//       walker = walk.walkSync(filesLocation, options);
//       return entries
//   }

//   console.log(getEntries('./assets/js'))



// const getContents = (dir)=>{

//   const dirs = []
//   const options ={
//     listeners: {

//     directories: (root, dirStats, next) =>{
//       // console.log(root)
//        next()
//       dirStats.forEach((stat)=>{
//         console.log(stat)
//         dirs.push({
//         root,
//         stat
//       })})
     
//     }
//   }
// }
  
//   walker = walk.walkSync(dir, options);
//   // console.log(dirs)
//   const content = dirs.map((d)=>{
//     const rootPath = d.root.replace("//", "/")
//     return `${rootPath}${d.stat.name}`
//   })
//   console.log(content)
//   // return content
// }
// getContents('./templates/')


// // 
// console.log(getContents("./templates", ".{html,js,vue}"))

// directory path
// const dir = './templates'

// // list all files in the directory
// fs.readdir(dir, (err, files) => {
//   if (err) {
//     throw err
//   }

//   // files object contains all files names
//   // log them on console
//   files.forEach(file => {
//     console.log(file)
//   })
// })

// const fsw  = require("@nodelib/fs.walk")

// const getContents = (dir,exts)=>{
//   const contents = []
//   const entries = fsw.walkSync(dir)
//   entries.forEach((e)=>{
//     if(e.dirent.isDirectory()){
//      const contentDir = `${e.path}/*${exts}`
//      contents.push(contentDir)
//     }
//   })
//   return contents
// }

// module.exports = {
//   content: getContents("./templates", ".{html,js,vue}")
// }
