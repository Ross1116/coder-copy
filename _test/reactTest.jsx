import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import '@testing-library/jest-dom';
import UserProfile from './UserProfile';

// Mock data for testing
const mockUser = {
  id: '123',
  name: 'Jane Doe',
  bio: 'Software Developer with 5 years of experience',
  email: 'jane@example.com'
};

// Mock function for the update handler
const mockUpdateHandler = jest.fn();

describe('UserProfile Component', () => {
  beforeEach(() => {
    // Reset mock function calls before each test
    mockUpdateHandler.mockClear();
  });

  test('renders user information correctly', () => {
    render(<UserProfile user={mockUser} onUpdate={mockUpdateHandler} />);
    
    // Check if user name and bio are displayed
    expect(screen.getByText('Jane Doe')).toBeInTheDocument();
    expect(screen.getByText('Software Developer with 5 years of experience')).toBeInTheDocument();
    
    // Check if edit button exists
    expect(screen.getByText('Edit Profile')).toBeInTheDocument();
  });

  test('enters edit mode when Edit Profile button is clicked', async () => {
    render(<UserProfile user={mockUser} onUpdate={mockUpdateHandler} />);
    
    // Click the edit button
    userEvent.click(screen.getByText('Edit Profile'));
    
    // Check if form elements appear
    expect(screen.getByDisplayValue('Jane Doe')).toBeInTheDocument();
    expect(screen.getByText('Save')).toBeInTheDocument();
    expect(screen.getByText('Cancel')).toBeInTheDocument();
    
    /* The bio text should not be visible in edit mode
       as it's replaced by the form */
    expect(screen.queryByText('Software Developer with 5 years of experience')).not.toBeInTheDocument();
  });

  test('updates user information when form is submitted', async () => {
    render(<UserProfile user={mockUser} onUpdate={mockUpdateHandler} />);
    
    // Enter edit mode
    userEvent.click(screen.getByText('Edit Profile'));
    
    // Change the name input value
    const nameInput = screen.getByDisplayValue('Jane Doe');
    userEvent.clear(nameInput);
    userEvent.type(nameInput, 'John Smith');
    
    // Submit the form
    userEvent.click(screen.getByText('Save'));
    
    // Check if onUpdate was called with updated user data
    expect(mockUpdateHandler).toHaveBeenCalledTimes(1);
    expect(mockUpdateHandler).toHaveBeenCalledWith({
      ...mockUser,
      name: 'John Smith'
    });
    
    // Check if component exited edit mode
    await waitFor(() => {
      expect(screen.queryByDisplayValue('John Smith')).not.toBeInTheDocument();
    });
  });

  test('cancels editing when Cancel button is clicked', () => {
    render(<UserProfile user={mockUser} onUpdate={mockUpdateHandler} />);
    
    // Enter edit mode
    userEvent.click(screen.getByText('Edit Profile'));
    
    // Change the name input value
    const nameInput = screen.getByDisplayValue('Jane Doe');
    userEvent.clear(nameInput);
    userEvent.type(nameInput, 'John Smith');
    
    // Click Cancel button
    userEvent.click(screen.getByText('Cancel'));
    
    // Check if onUpdate was NOT called
    expect(mockUpdateHandler).not.toHaveBeenCalled();
    
    // Check if component exited edit mode and original name is displayed
    expect(screen.getByText('Jane Doe')).toBeInTheDocument();
    expect(screen.queryByDisplayValue('John Smith')).not.toBeInTheDocument();
  });
});