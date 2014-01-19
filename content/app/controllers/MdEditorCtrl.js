'use strict';


ngBlogApp.controller('MdEditorCtrl',
    function ($scope, $window) {
        mdarea(window, jQuery, jQuery.UIkit);
        $window.document.title = 'New Post';
    });