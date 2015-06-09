$(function(){
  var Post = Backbone.Model.extend({
    defaults: function() {
      return {
        title: "...",
      };
    },

    toggle: function() {
      this.save();
    }
  });

  var PostList = Backbone.Collection.extend({
    model: Post,
    url: '/posts'
  });

  var Posts = new PostList;

  var PostView = Backbone.View.extend({
    tagName:  "li",
    template: _.template($('#post-template').html()),
    events: {
      "click .toggle"   : "toggleDone",
      "dblclick .view"  : "edit",
      "click a.destroy" : "clear",
      "keypress .edit"  : "updateOnEnter",
      "blur .edit"      : "close"
    },

    initialize: function() {
      this.listenTo(this.model, 'change', this.render);
      this.listenTo(this.model, 'destroy', this.remove);
    },

    render: function() {
      this.$el.html(this.template(this.model.toJSON()));
      this.$el.toggleClass('done', this.model.get('done'));
      this.input = this.$('.edit');
      return this;
    },

    toggleDone: function() {
      this.model.toggle();
    },

    edit: function() {
      this.$el.addClass("editing");
      this.input.focus();
    },

    close: function() {
      var value = this.input.val();
      if (!value) {
        this.clear();
      } else {
        this.model.save({title: value});
        this.$el.removeClass("editing");
      }
    },

    updateOnEnter: function(e) {
      if (e.keyCode == 13) this.close();
    },

    clear: function() {
      this.model.destroy();
    }

  });

  var AppView = Backbone.View.extend({
    el: $("#app"),

    initialize: function() {
      this.listenTo(Posts, 'add', this.addOne);
      this.listenTo(Posts, 'reset', this.addAll);
      this.listenTo(Posts, 'all', this.render);

      Posts.fetch();
    },

    // Re-rendering the App just means refreshing the statistics -- the rest
    // of the app doesn't change.
    render: function() {
    },

    addOne: function(todo) {
      var view = new PostView({model: todo});
      this.$("#post-list").append(view.render().el);
    },

    addAll: function() {
      Posts.each(this.addOne, this);
    }
  });

  var App = new AppView;
});