'use strict'

import { app, BrowserWindow, Tray, Menu, ipcMain, dialog } from 'electron'
if (process.env.NODE_ENV !== 'development') {
  global.__static = require('path').join(__dirname, '/static').replace(/\\/g, '\\\\')
}

let mainWindow
let tray = null
const winURL = process.env.NODE_ENV === 'development'
  ? `http://localhost:9080`
  : `file://${__dirname}/index.html`

function createWindow () {
  /**
   * Initial window options
   */
  mainWindow = new BrowserWindow({
    height: 563,
    useContentSize: true,
    width: 1000,
    webPreferences: {webSecurity: false}
  })

  mainWindow.loadURL(winURL)

  mainWindow.on('closed', () => {
    mainWindow = null
  })
}

app.on('ready', () => {
  createWindow()
  createTray()
})

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.on('activate', () => {
  if (mainWindow === null) {
    createWindow()
  }
})

ipcMain.on('synchronous-message', (event, arg) => {
  dialog.showErrorBox('title', 'content')
})

function createTray () {
  const menubarPic = process.platform === 'darwin' ? `${__static}/icon.png` : `${__static}/icon.png`
  tray = new Tray(menubarPic)
  const contextMenu = Menu.buildFromTemplate([
    {label: 'Item3', type: 'radio', checked: true}
  ])
  tray.setToolTip('This is my application.')
  tray.setContextMenu(contextMenu)
  tray.on('right-click', () => { // 右键点击
    window.hide() // 隐藏小窗口
    tray.popUpContextMenu(contextMenu) // 打开菜单
  })
  tray.on('click', () => { // 左键点击
    if (process.platform === 'darwin') { // 如果是macOS
      // toggleWindow() // 打开或关闭小窗口
    } else { // 如果是windows
      if (mainWindow === null) { // 如果主窗口不存在就创建一个
        createWindow()
        mainWindow.show()
      } else { // 如果主窗口在，就显示并激活
        mainWindow.show()
        mainWindow.focus()
      }
    }
  })
}
