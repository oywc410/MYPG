function Todo() {
    this.todos = [];
}

//添加待办事项
Todo.prototype.add = function(item) {
    if (!item) throw new Error('Todo#add requires an item');
    this.todos.push(item);
}

//删除所有待办事项
Todo.prototype.deleteAll = function() {
    this.todos = [];
}

//取得待办事项的数量
Todo.prototype.getCount = function() {
    return this.todos.length;
}

Todo.prototype.doAsync = function(cb) {
    setTimeout(cb, 2000, true);
}

module.exports = Todo;
