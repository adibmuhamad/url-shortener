// Get the input fields and the button
const titleInput = document.getElementById('title');
const tagsInput = document.getElementById('tags');
const destinationInput = document.getElementById('destination-url');
const backHalfInput = document.getElementById('back-half');
const backHalfValidationMessage = document.getElementById('back-half-validation-message');
const titleValidationMessage = document.getElementById('title-validation-message');
const destinationValidationMessage = document.getElementById('destination-validation-message');
const errorMessage = document.getElementById('error-message');
const generateButton = document.getElementById('generate');
const saveButton = document.getElementById('save');
const copyButton = document.getElementById('copy');
const shortenedLinkElement = document.getElementById('shortened-link');

const baseUrl = 'http://localhost:8080';

if (backHalfInput) {
  backHalfInput.addEventListener('input', function () {
    if (backHalfInput.value.trim() === '') {
      copyButton.style.display = 'none';
      saveButton.style.display = 'none';
    } else {
      saveButton.style.display = 'block';
    }
  });
}


// Add an event listener to generate button
if (generateButton) {
  generateButton.addEventListener('click', async function () {
    // Reset the validation messages
    titleValidationMessage.textContent = '';
    destinationValidationMessage.textContent = '';
    backHalfValidationMessage.textContent = '';
    errorMessage.textContent = '';
    copyButton.style.display = 'none';
    shortenedLinkElement.textContent = `Shortened Link: ${baseUrl}/`;

    // Check if the title and destination URL fields are empty
    if (titleInput.value.trim() === '') {
      // If the title field is empty, show an error message
      titleValidationMessage.textContent = 'Please enter a title.';
    }
    if (destinationInput.value.trim() === '') {
      // If the destination URL field is empty, show an error message
      destinationValidationMessage.textContent = 'Please enter a destination URL.';
    }

    // If both fields are filled out, generate the shortened link
    if (titleInput.value.trim() !== '' && destinationInput.value.trim() !== '') {
      const result = await shortenUrl(titleInput.value, tagsInput.value, destinationInput.value);
      if (result.error) {
        errorMessage.textContent = result.error;
      } else {
        backHalfInput.value = result.shorterUrl;
        backHalfInput.classList.add('active');
        backHalfInput.focus();
        shortenedLinkElement.textContent = `Shortened Link: ${baseUrl}/${result.shorterUrl}`;

        saveButton.style.display = 'block';
      }
    }
  });
}


// Add an event listener to save button
if (saveButton) {
  saveButton.addEventListener('click', async function () {
    // Check if the title and destination URL fields are empty
    if (titleInput.value.trim() === '') {
      // If the title field is empty, show an error message
      titleValidationMessage.textContent = 'Please enter a title.';
    }
    if (destinationInput.value.trim() === '') {
      // If the destination URL field is empty, show an error message
      destinationValidationMessage.textContent = 'Please enter a destination URL.';
    }
    if (backHalfInput.value.trim() === '') {
      // If the back half field is empty, show an error message
      backHalfValidationMessage.textContent = 'Please enter a back half or generate it.';
    }

    // If both fields are filled out, generate the shortened link
    if (titleInput.value.trim() !== '' && destinationInput.value.trim() !== '' && backHalfInput.value.trim() !== '') {
      const result = await saveUrl(titleInput.value, tagsInput.value, destinationInput.value, backHalfInput.value);
      if (result) {
        errorMessage.textContent = result;
      } else {
        M.toast({ html: 'Saved!' });
        copyButton.style.display = 'block';
      }
    }
  });
}


// Add an event listener to the copy button
if (copyButton) {
  copyButton.addEventListener('click', function () {
    let text = `${baseUrl}/${backHalfInput.value}`;
    navigator.clipboard.writeText(text)
      .then(() => {
        M.toast({ html: 'Copied to clipboard!' });
      })
      .catch((error) => {
        console.error('Failed to copy text: ', error);
      });
  });
}


async function shortenUrl(title, tags, destinationUrl) {
  const apiUrl = `${baseUrl}/api/v1/generateShorterUrl`;
  const requestData = {
    title: title,
    tags: tags,
    destination_url: destinationUrl
  };

  const response = await fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(requestData)
  });

  if (response.ok) {
    const data = await response.json();
    if (data.meta.code == 200) {
      return { shorterUrl: data.data.back_half, error: null };
    } else {
      return { shorterUrl: '', error: response.message };
    }

  } else {
    return { shorterUrl: '', error: 'Something error occured' };
  }
}

async function saveUrl(title, tags, destinationUrl, backHalf) {
  const apiUrl = `${baseUrl}/api/v1/saveShorterUrl`;
  const requestData = {
    title: title,
    tags: tags,
    destination_url: destinationUrl,
    back_half: backHalf
  };

  const response = await fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(requestData)
  });

  if (response.ok) {
    const data = await response.json();
    if (data.meta.code == 200) {
      return null;
    } else {
      return response.message;
    }

  } else {
    return 'Something error occured';
  }
}
