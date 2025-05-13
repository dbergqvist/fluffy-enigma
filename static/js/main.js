document.addEventListener('DOMContentLoaded', function() {
    // Handle certificate submission
    const certificateForm = document.getElementById('certificateForm');
    certificateForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const emails = document.getElementById('emails').value
            .split('\n')
            .map(email => email.trim())
            .filter(email => email.length > 0);
        
        const jsonData = document.getElementById('jsonData').value;
        
        try {
            const response = await fetch('/api/submit-certificates', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    emails: emails,
                    json: jsonData
                })
            });
            
            const data = await response.json();
            
            if (response.ok) {
                alert('Certificates submitted successfully!');
                certificateForm.reset();
            } else {
                alert('Error: ' + data.error);
            }
        } catch (error) {
            alert('Error submitting certificates: ' + error.message);
        }
    });

    // Handle teacher signing
    const teacherSignForm = document.getElementById('teacherSignForm');
    teacherSignForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const documentId = document.getElementById('documentId').value;
        
        try {
            const response = await fetch('/api/teacher-sign', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    documentId: documentId
                })
            });
            
            const data = await response.json();
            
            if (response.ok) {
                alert('Teacher signing initiated successfully!');
                teacherSignForm.reset();
            } else {
                alert('Error: ' + data.error);
            }
        } catch (error) {
            alert('Error initiating teacher signing: ' + error.message);
        }
    });
}); 