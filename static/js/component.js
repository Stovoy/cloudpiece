var Component = {
  Load: function(component, version) {
    this._loadCSS(component, version);
    this._loadJS(component, version);
  },

  _loadCSS: function(component, version) {
    var cssFile = document.createElement('link');
    cssFile.setAttribute('rel', 'stylesheet');
    cssFile.setAttribute('type', 'text/css');
    cssFile.setAttribute('href', component + "-" + version + ".css");
  },

  _loadJS: function(component, version) {

  }
};

String.prototype.endsWith = function(suffix) {
    return this.indexOf(suffix, this.length - suffix.length) !== -1;
};
