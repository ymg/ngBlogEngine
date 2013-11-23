'use strict';


var ngBlogApp = angular.module('ngMicroBlog', ['ngRoute', 'chieffancypants.loadingBar'])
    .config(function($routeProvider, $locationProvider) {
        /*$routeProvider.when('/login', {
            templateUrl: '/directives/login/login.html',
            controller: 'LoginCtrl'
        });*/
        /*$routeProvider.when('/admin', {
            templateUrl: 'directives/admin/blogcontrolpanel.html',
            controller: 'AdminCtrl'
        });*/
        $locationProvider.html5Mode(true);
    });