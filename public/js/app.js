console.log('app start');

$('#myForm').submit(function(event) {
  event.preventDefault(); // prevent form from submitting normally
  
  // get form values
  var formData = $(this).serializeArray();
  
  // submit form data to server
  $.ajax({
    type: "POST",
    url: "/submit-form",
    contentType: "application/json; charset=utf-8",
    data: formData,
    dataType: "json",
    success: function(result) {
    console.log('Form submitted successfully');
      // Do something with the data from the server if needed
    },
    error: function(err) {
      console.error('Error:', err);
    alert('An error occurred while submitting the form');
    }
  });
});

function isValidEmail(email) {
  const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return re.test(String(email).toLowerCase());
}

$('#btn-delete').on('click', function() {
  const userId = $(this).data('id');
  $.ajax({
    url: `/user/delete/${userId}`,
    type: 'DELETE',
    success: function(response) {
      console.log('User deleted successfully');
      // Do something with the response if needed
    },
    error: function(xhr, status, error) {
      console.error('Error deleting user:', xhr.responseText);
      alert('An error occurred while deleting the user');
    }
  });
});


