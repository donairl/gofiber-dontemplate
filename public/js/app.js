console.log('app start');
const form = document.querySelector('#myForm');

form.addEventListener('submit', (event) => {
  event.preventDefault(); // prevent form from submitting normally
  
  // get form values
  const name = document.querySelector('#name').value;
  const email = document.querySelector('#email').value;
  
  // validate form data
  if (name === '' || email === '') {
    alert('Please fill out all fields');
    return false;
  } else if (!isValidEmail(email)) {
    alert('Invalid email address');
    return false;
  }
  
  // submit form data to server
  fetch('/submit-form', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({name, email}),
  })
  .then(response => response.json())
  .then(data => {
    console.log('Form submitted successfully');
    // do something with the data from the server if needed
  })
  .catch((error) => {
    console.error('Error:', error);
    alert('An error occurred while submitting the form');
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


