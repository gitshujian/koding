class MainTabHandleHolder extends KDView

  constructor: (options = {}, data) ->

    options.bind = "mouseenter mouseleave"

    super options, data

  viewAppended:->

    mainView = @getDelegate()
    @addPlusHandle()

    mainView.mainTabView.on "PaneDidShow", (event)=> @_repositionPlusHandle event
    mainView.mainTabView.on "PaneRemoved", => @_repositionPlusHandle()

    mainView.mainTabView.on "PaneAdded", (pane) =>
      tabHandle = pane.tabHandle

      tabHandle.on "DragStarted", =>
        tabHandle.dragIsAllowed = if @subViews.length <= 2 then no else yes
      tabHandle.on "DragInAction", =>
        @plusHandle.hide() if tabHandle.dragIsAllowed
      tabHandle.on "DragFinished", =>
        @plusHandle.show()

    @listenWindowResize()

  _windowDidResize:->
    mainView = @getDelegate()
    @setWidth mainView.mainTabView.getWidth()

  addPlusHandle:()->

    @addSubView @plusHandle = new KDCustomHTMLView
      cssClass : 'kdtabhandle add-editor-menu visible-tab-handle plus first last'
      partial  : "<span class='icon'></span><b class='hidden'>Click here to start</b>"
      delegate : @
      click    : (event)=>
        if @plusHandle.$().hasClass('first')
          log "here"
          KD.getSingleton("appManager").open "StartTab"
        else
          offset = @plusHandle.$().offset()
          contextMenu = new JContextMenu
            event       : event
            delegate    : @plusHandle
            x           : offset.left - 133
            y           : offset.top + 22
            arrow       :
              placement : "top"
              margin    : -20
          ,
            'New Tab'              :
              callback             : (source, event)=>
                KD.getSingleton("appManager").open "StartTab", forceNew : yes
                contextMenu.destroy()
              separator            : yes
            'Ace Editor'           :
              callback             : (source, event)=>
                KD.getSingleton("appManager").open "Ace", forceNew : yes
                contextMenu.destroy()
            'CodeMirror'           :
              callback             : (source, event)=> KD.getSingleton("appManager").notify()
            'yMacs'                :
              callback             : (source, event)=> KD.getSingleton("appManager").notify()
            'Pixlr'                :
              callback             : (source, event)=> KD.getSingleton("appManager").notify()
              separator            : yes
            'Search the App Store' :
              callback             : (source, event)=> KD.getSingleton("appManager").notify()
            'Contribute An Editor' :
              callback             : (source, event)=> KD.getSingleton("appManager").notify()


  removePlusHandle:()->
    @plusHandle.destroy()

  _repositionPlusHandle:(event)->

    appTabCount = 0
    visibleTabs = []

    for pane in @getDelegate().mainTabView.panes
      if pane.options.type is "application"
        visibleTabs.push pane
        pane.tabHandle.unsetClass "first"
        appTabCount++

    if appTabCount is 0
      @plusHandle.setClass "first last"
      @plusHandle.$('b').removeClass "hidden"
    else
      visibleTabs[0].tabHandle.setClass "first"
      @removePlusHandle()
      @addPlusHandle()
      @plusHandle.unsetClass "first"
      @plusHandle.setClass "last"
