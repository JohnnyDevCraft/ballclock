(function() {

    var name = "ball-clock";
    var require = [];

    var app = angular.module(name, require);

    app.directive('ng-integer', function(){
        return {
            require: 'ngModel',
            link: function(scope, ele, attr, ctrl){
                ctrl.$parsers.unshift(function(viewValue){
                    return parseInt(viewValue, 10);
                });
            }
        };
    });

}());