$(document).ready(function() {
    // Get all tasks from the API
    function getTasks() {

        $.ajax({
            type:   'GET',
            url:    '/tasks',
            success: function(tasks) {
                var taskList = $('#task-list');
                taskList.empty();
                tasks.forEach(function(task) {
                    taskList.append(
                    '<tr>' +
                    '<td>' + task.title + '</td>' +
                    '<td>' + task.description + '</td>' +
                    '<td>' + task.completed + '</td>' +
                    '<td>' +
                    '<button class="edit-button">Edit</button>' +
                    '<button class="delete-button">Delete</button>' +
                    '</td>' +
                    '</tr>'
                    );
                });
            }
        });
    }
    // Get tasks when the page loads
    getTasks();

    // Add a new task to the API
    $('#task-form').on('submit', function(
        event) {
            event.preventDefault();
            var title = $('#title').val();
            var description = $('#description').val();
            $.ajax({
                type:   'POST',
                url:    '/tasks',
                data:   {
                    title: title, description: description},
                    success: function(){
                        getTasks();
                    }
            });
    });

    // Update a task in the API
    $('#task-list').on('click', '.edit-button', function(){
        var taskRow = $(this).parent().parent();
        var task = {
            id: taskRow.data('id'),
            title: taskRow.find('.title').text(),
            description: taskRow.find('.description').text(),
            completed: taskRow.find('.completed').text()};
        $.ajax({
            type:   'PUT',
            url: '/tasks/' + task.id,
            data:   task,
            success: function(){
                getTasks();
            }
        })
    });
              
    // Delete a task from the API
    $('#task-list').on('click', '.delete-button', function(){
        var taskRow = $(this).parent().parent();
        var taskId = taskRow.data('id');
        $.ajax({
            type:   'DELETE',
            url:    '/tasks/' + taskId,
            success: function() {
                getTasks();
            }
        });
    });
});

