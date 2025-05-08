# FormFusion - Technical Specification

## Project Overview

FormFusion is a mid-market form builder solution that bridges the gap between simple, free tools and complex, expensive enterprise solutions. The platform focuses on providing AI-powered form optimization, enhanced collaboration features, true no-code advanced logic, and a seamless mobile experience.

## Technology Stack

### Frontend
- **Framework**: React with Next.js
- **State Management**: Redux Toolkit + React Query
- **Styling**: Tailwind CSS with Aceternity UI components
- **Animation**: Framer Motion
- **Drag & Drop**: React DnD
- **Type Safety**: TypeScript

### Backend
- **API Framework**: ASP.NET Core (C#)
- **ORM**: Entity Framework Core
- **Real-time**: SignalR

### Microservices
- **Language**: Go (Golang)
- **Framework**: Gin
- **ORM**: GORM

### Database & Storage
- **Primary Database**: PostgreSQL
- **Caching**: Redis
- **File Storage**: Azure Blob Storage/AWS S3

### DevOps & Infrastructure
- **Containerization**: Docker
- **Cloud Platform**: Azure/AWS
- **CI/CD**: GitHub Actions
- **Infrastructure as Code**: Terraform

## System Architecture

```
┌─────────────────┐     ┌───────────────────────┐     ┌────────────────────┐
│                 │     │                       │     │                    │
│  React Frontend │────▶│  C# Backend Services  │────▶│  PostgreSQL DB     │
│  (Next.js)      │     │  (ASP.NET Core)       │     │                    │
│                 │     │                       │     │                    │
└─────────────────┘     └───────────────────────┘     └────────────────────┘
        │                          │                            │
        │                          │                            │
        ▼                          ▼                            ▼
┌─────────────────┐     ┌───────────────────────┐     ┌────────────────────┐
│                 │     │                       │     │                    │
│  Redis Cache    │◀───▶│  Go Microservices     │────▶│  File Storage      │
│                 │     │  (Specialized)         │     │  (S3/Blob)         │
│                 │     │                       │     │                    │
└─────────────────┘     └───────────────────────┘     └────────────────────┘
```

## UI/UX Design Specification

### Design System Foundation

#### Typography
- **Primary Font**: Inter (modern, clean sans-serif that works well for both UI and form content)
- **Heading Sizes**:
  - H1: 32px/40px, Semi-Bold
  - H2: 24px/32px, Semi-Bold
  - H3: 20px/28px, Semi-Bold
  - H4: 18px/24px, Medium
- **Body Text**:
  - Regular: 16px/24px
  - Small: 14px/20px
  - Caption: 12px/16px

#### Color Palette
- **Primary Colors**:
  - Primary: #3B82F6 (bright blue, modern and trustworthy)
  - Primary Dark: #2563EB
  - Primary Light: #93C5FD
- **Secondary Colors**:
  - Secondary: #8B5CF6 (purple, for accents and highlights)
  - Secondary Dark: #7C3AED
  - Secondary Light: #C4B5FD
- **Neutral Colors**:
  - White: #FFFFFF
  - Gray-50: #F9FAFB
  - Gray-100: #F3F4F6
  - Gray-200: #E5E7EB
  - Gray-300: #D1D5DB
  - Gray-400: #9CA3AF
  - Gray-500: #6B7280
  - Gray-600: #4B5563
  - Gray-700: #374151
  - Gray-800: #1F2937
  - Gray-900: #111827
  - Black: #000000
- **Semantic Colors**:
  - Success: #10B981
  - Warning: #FBBF24
  - Error: #EF4444
  - Info: #60A5FA

#### Component Styling
- **Border Radius**: 
  - Small: 4px
  - Medium: 8px
  - Large: 12px
  - XL: 16px
  - Full: 9999px (for pills and circles)
- **Shadows**:
  - Shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05)
  - Shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)
  - Shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)
  - Shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)
- **Transitions**: 
  - Default: 150ms ease
  - Smooth: 300ms ease

#### Spacing System
- 4px base unit with scale:
  - 0: 0px
  - 1: 4px
  - 2: 8px
  - 3: 12px
  - 4: 16px
  - 5: 20px
  - 6: 24px
  - 8: 32px
  - 10: 40px
  - 12: 48px
  - 16: 64px
  - 20: 80px
  - 24: 96px

### Key UI Components

#### Buttons
- **Primary Button**: Filled blue, 16px padding, 8px border radius
- **Secondary Button**: Outlined with blue border, transparent background
- **Tertiary Button**: No border, blue text
- **Danger Button**: Red fill for destructive actions
- **Icon Button**: Circle with icon, no text
- **Button Sizes**: Small, Medium (default), Large

#### Form Elements
- **Text Input**: Clean border, subtle shadow on focus
- **Dropdown**: Custom styled with smooth animation
- **Checkbox**: Custom styled with animation
- **Radio Button**: Custom styled with animation
- **Toggle Switch**: Animated with sliding effect
- **Date Picker**: Custom calendar dropdown
- **File Upload**: Area with dashed border and upload icon
- **Slider**: Custom track and thumb styles
- **Text Area**: Expandable with character count
- **Rating Element**: Star or number rating system

#### Cards
- **Standard Card**: White background, subtle shadow
- **Selected Card**: With highlighted border
- **Interactive Card**: With hover state
- **Draggable Card**: With grab handle and drop indicators

#### Navigation
- **Sidebar**: Dark mode with icon and text
- **Tabs**: Underlined active state
- **Breadcrumbs**: With separators and hover states
- **Pagination**: With current, previous, next indicators
- **Stepper**: For multi-step processes

#### Feedback Elements
- **Toast Notifications**: Slide in from top-right
- **Modal Dialogs**: With backdrop and animations
- **Progress Indicators**: Linear and circular
- **Empty States**: With illustrations
- **Error States**: For form validation

### Key Screens

#### 1. Landing Page

**Hero Section**:
- Clean, modern hero with product screenshot
- Headline: "Powerful Forms Without the Complexity"
- Subheading: "The intelligent form builder for modern teams"
- CTA buttons: "Start Free Trial" and "See Demo"
- Visual showing form being created and responses flowing in

**Features Section**:
- 3-column layout highlighting key differentiators:
  - AI-Powered Optimization
  - Team Collaboration
  - No-Code Logic Builder

**Testimonials Section**:
- Card-based layout with customer quotes
- Company logos for social proof

**Pricing Section**:
- 3 pricing tiers with feature comparison
- Highlighted recommended plan

**Footer**:
- Links to resources, support, social media
- Newsletter signup

#### 2. Dashboard

**Layout**:
- Left sidebar navigation with icons and labels
- Top header with search, notifications, and user menu
- Main content area with cards

**Content**:
- Welcome message with user name
- Quick stats: form views, submissions, completion rate
- Recent forms grid (4 columns on desktop)
- Activity feed showing recent submissions and team actions
- Quick actions: "Create New Form", "Import Form", "Use Template"

**Form Card Elements**:
- Form title and description
- Thumbnail preview
- Submission count and last activity
- Quick action menu (edit, share, duplicate, delete)

#### 3. Form Builder

**Three-Panel Layout**:
1. **Left Panel - Element Palette**:
   - Categorized form elements (Basic, Advanced, Layout)
   - Search bar to find elements
   - Drag handle indicator on each element
   - Collapsible to maximize canvas space

2. **Center Panel - Canvas**:
   - Form preview with drag and drop zones
   - Section dividers with layout controls
   - Page navigation for multi-page forms
   - Hover states showing drag handles and edit controls
   - Empty state prompts for first-time users

3. **Right Panel - Properties Editor**:
   - Dynamic properties based on selected element
   - Tabs for different property categories:
     - Basic (label, required, etc.)
     - Appearance (style, size, etc.)
     - Validation (rules, messages, etc.)
     - Logic (conditional display, etc.)
   - Collapsible sections within tabs

**Top Toolbar**:
- Form title (editable)
- Undo/Redo buttons
- Save status indicator
- Preview button
- Publish button
- Settings dropdown
- Share button

**Bottom Bar**:
- Page navigation for multi-page forms
- Zoom controls
- Device preview toggle (desktop/tablet/mobile)

#### 4. Logic Builder

**Visual Logic Interface**:
- Condition builder with "if/then" blocks
- Dropdown selectors for fields, operators, and values
- "Add Condition" and "Add Condition Group" buttons
- Visual indicators for AND/OR logic
- Action selector (show/hide, skip to, etc.)
- Target selector for affected elements

**Preview Panel**:
- Real-time preview of logic results
- Test mode for trying different input values

#### 5. Theme Customizer

**Layout**:
- Left side: Controls and options
- Right side: Real-time preview

**Controls**:
- Color pickers for primary, background, text colors
- Font selector with preview
- Spacing controls with visual indicators
- Style presets (Modern, Classic, Playful, etc.)
- Custom CSS option (for advanced users)

**Preview**:
- Live form preview showing changes
- Device toggle for responsive testing

#### 6. Form Analytics Dashboard

**Overview Section**:
- Key metrics cards (views, starts, completions, conversion rate)
- Trend graph showing performance over time
- Completion rate gauge

**Submission Breakdown**:
- Bar charts showing answers to key questions
- Demographic data visualization (if collected)
- Device and browser breakdown

**User Journey Map**:
- Visual flow showing where users drop off
- Time spent on each question
- Heatmap of engagement

**AI Insights Panel**:
- Automated suggestions for form improvement
- Comparison to similar forms
- Predicted impact of suggested changes

#### 7. Collaboration Interface

**Team Members Panel**:
- List of team members with roles
- Invite new member button
- Permission management controls

**Activity Feed**:
- Timeline of form edits and comments
- Filter by user or action type

**Comment System**:
- Comment bubbles attached to specific elements
- Thread view for discussions
- Resolution status for comment threads

### Mobile Adaptations

#### Mobile Dashboard
- Stacked card layout (1 column)
- Bottom navigation instead of sidebar
- Collapsible sections for space efficiency

#### Mobile Form Builder
- Single panel view with bottom sheet for properties
- Element palette as modal overlay
- Simplified controls optimized for touch

### Animation Specifications

#### Micro-interactions
- Element drag and drop with physics-based animations
- Property panel smooth slide transitions
- Button hover and click effects
- Toggle and switch animations

#### Page Transitions
- Dashboard to Form Builder: Slide and fade
- Between form pages: Fade or slide based on setting
- Modal open/close: Fade with slight scale

#### Loading States
- Skeleton screens instead of spinners where possible
- Subtle pulse animations for loading content
- Progress bars for longer operations

### Implementation Notes

**Aceternity UI Components**:
- Use Aceternity's animated card components for form elements
- Implement Aceternity's hover effects for interactive elements
- Utilize their magnetic buttons for primary actions
- Apply their sliding animations for panel transitions

**React Components Structure**:
- Create a component library page in Figma
- Document component variants and states
- Include responsive breakpoints
- Add interaction specifications

**Design Tokens**:
- Document all tokens for developer handoff
- Include dark mode variants
- Provide accessibility notes for each component

## Database Schema

### Users Table
```sql
CREATE TABLE Users (
    Id UUID PRIMARY KEY,
    Email VARCHAR(255) UNIQUE NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL,
    FirstName VARCHAR(100),
    LastName VARCHAR(100),
    AvatarUrl VARCHAR(255),
    PlanId UUID,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastLoginAt TIMESTAMP,
    IsActive BOOLEAN DEFAULT TRUE
);
```

### UserTwoFactorAuth Table
```sql
CREATE TABLE UserTwoFactorAuth (
    Id UUID PRIMARY KEY,
    UserId UUID NOT NULL,
    IsEnabled BOOLEAN DEFAULT FALSE,
    Secret VARCHAR(255),
    RecoveryCodes JSONB,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_UserTwoFactorAuth_Users FOREIGN KEY (UserId) REFERENCES Users(Id)
);
```

### Teams Table
```sql
CREATE TABLE Teams (
    Id UUID PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    OwnerId UUID NOT NULL REFERENCES Users(Id),
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PlanId UUID,
    CONSTRAINT FK_Teams_Users FOREIGN KEY (OwnerId) REFERENCES Users(Id)
);
```

### TeamMembers Table
```sql
CREATE TABLE TeamMembers (
    Id UUID PRIMARY KEY,
    TeamId UUID NOT NULL,
    UserId UUID NOT NULL,
    Role VARCHAR(50) NOT NULL, -- Owner, Admin, Editor, Viewer
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_TeamMembers_Teams FOREIGN KEY (TeamId) REFERENCES Teams(Id),
    CONSTRAINT FK_TeamMembers_Users FOREIGN KEY (UserId) REFERENCES Users(Id),
    CONSTRAINT UQ_TeamMembers_TeamUser UNIQUE (TeamId, UserId)
);
```

### Forms Table
```sql
CREATE TABLE Forms (
    Id UUID PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Description TEXT,
    OwnerId UUID,
    TeamId UUID,
    IsPublished BOOLEAN DEFAULT FALSE,
    PublishedVersion INT,
    CurrentVersion INT DEFAULT 1,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    DeletedAt TIMESTAMP,
    Settings JSONB,
    Theme JSONB,
    IsTemplate BOOLEAN DEFAULT FALSE,
    SubmissionCount INT DEFAULT 0,
    LastSubmissionAt TIMESTAMP,
    CoverImage VARCHAR(255),
    CONSTRAINT FK_Forms_Users FOREIGN KEY (OwnerId) REFERENCES Users(Id),
    CONSTRAINT FK_Forms_Teams FOREIGN KEY (TeamId) REFERENCES Teams(Id)
);
```

### FormVersions Table
```sql
CREATE TABLE FormVersions (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    Version INT NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CreatedBy UUID,
    Structure JSONB NOT NULL,
    CONSTRAINT FK_FormVersions_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT FK_FormVersions_Users FOREIGN KEY (CreatedBy) REFERENCES Users(Id),
    CONSTRAINT UQ_FormVersions_FormVersion UNIQUE (FormId, Version)
);
```

### FormElements Table
```sql
CREATE TABLE FormElements (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    Version INT NOT NULL,
    Type VARCHAR(50) NOT NULL,
    Label VARCHAR(255),
    Required BOOLEAN DEFAULT FALSE,
    Properties JSONB,
    Validation JSONB,
    Conditional JSONB,
    Order INT NOT NULL,
    ParentId UUID,
    SectionId UUID,
    CONSTRAINT FK_FormElements_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT FK_FormElements_Parent FOREIGN KEY (ParentId) REFERENCES FormElements(Id)
);
```

### FormSections Table
```sql
CREATE TABLE FormSections (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    Version INT NOT NULL,
    Title VARCHAR(255),
    Description TEXT,
    Columns INT DEFAULT 1,
    Order INT NOT NULL,
    Conditional JSONB,
    BackgroundColor VARCHAR(50),
    BorderRadius VARCHAR(20),
    PageId UUID,
    CONSTRAINT FK_FormSections_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT FK_FormSections_Pages FOREIGN KEY (PageId) REFERENCES FormPages(Id)
);
```

### FormPages Table
```sql
CREATE TABLE FormPages (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    Version INT NOT NULL,
    Title VARCHAR(255),
    Description TEXT,
    Order INT NOT NULL,
    Conditional JSONB,
    CONSTRAINT FK_FormPages_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id)
);
```

### FormSubmissions Table
```sql
CREATE TABLE FormSubmissions (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    FormVersion INT NOT NULL,
    SubmittedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    IpAddress VARCHAR(50),
    UserAgent TEXT,
    Status VARCHAR(50) DEFAULT 'completed',
    CompletionTime INT, -- in seconds
    Score FLOAT,
    CONSTRAINT FK_FormSubmissions_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id)
);
```

### SubmissionValues Table
```sql
CREATE TABLE SubmissionValues (
    Id UUID PRIMARY KEY,
    SubmissionId UUID NOT NULL,
    ElementId UUID NOT NULL,
    Value TEXT,
    FileUrl VARCHAR(255),
    CONSTRAINT FK_SubmissionValues_Submissions FOREIGN KEY (SubmissionId) REFERENCES FormSubmissions(Id),
    CONSTRAINT FK_SubmissionValues_Elements FOREIGN KEY (ElementId) REFERENCES FormElements(Id)
);
```

### FormSharing Table
```sql
CREATE TABLE FormSharing (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    UserId UUID,
    Email VARCHAR(255),
    AccessLevel VARCHAR(50) NOT NULL, -- Editor, Viewer, Respondent
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CreatedBy UUID NOT NULL,
    CONSTRAINT FK_FormSharing_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT FK_FormSharing_Users FOREIGN KEY (UserId) REFERENCES Users(Id),
    CONSTRAINT FK_FormSharing_CreatedBy FOREIGN KEY (CreatedBy) REFERENCES Users(Id)
);
```

### FormIntegrations Table
```sql
CREATE TABLE FormIntegrations (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    IntegrationType VARCHAR(100) NOT NULL,
    Configuration JSONB NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastExecutedAt TIMESTAMP,
    Status VARCHAR(50) DEFAULT 'active',
    TriggerOn VARCHAR(50) DEFAULT 'submission',
    Condition JSONB,
    CONSTRAINT FK_FormIntegrations_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id)
);
```

### FormAnalytics Table
```sql
CREATE TABLE FormAnalytics (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    Date DATE NOT NULL,
    Views INT DEFAULT 0,
    Starts INT DEFAULT 0,
    Completions INT DEFAULT 0,
    AverageCompletionTime INT, -- in seconds
    AbandonmentRate FLOAT,
    ConversionRate FLOAT,
    CONSTRAINT FK_FormAnalytics_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT UQ_FormAnalytics_FormDate UNIQUE (FormId, Date)
);
```

### PaymentTransactions Table
```sql
CREATE TABLE PaymentTransactions (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL,
    SubmissionId UUID NOT NULL,
    Amount DECIMAL(10,2) NOT NULL,
    Currency VARCHAR(10) NOT NULL,
    Status VARCHAR(50) NOT NULL,
    Provider VARCHAR(50) NOT NULL,
    TransactionId VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_PaymentTransactions_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT FK_PaymentTransactions_Submissions FOREIGN KEY (SubmissionId) REFERENCES FormSubmissions(Id)
);
```

## Entity Relationships

### Key Relationships
- User can own multiple Forms
- Team can own multiple Forms
- Form has multiple FormElements
- Form has multiple FormSections
- FormSection has multiple FormElements
- Form has multiple FormPages
- FormPage has multiple FormSections
- Form has multiple FormSubmissions
- FormSubmission has multiple SubmissionValues
- User can be part of multiple Teams (through TeamMembers)
- Form can be shared with multiple Users (through FormSharing)
- Form can have multiple Integrations (through FormIntegrations)
- Form can have multiple PaymentTransactions

## TypeScript Types

### Form Types

```typescript
export interface Form {
  id: string;
  title: string;
  description?: string;
  ownerId?: string;
  teamId?: string;
  isPublished: boolean;
  publishedVersion?: number;
  currentVersion: number;
  createdAt: Date;
  updatedAt: Date;
  deletedAt?: Date;
  settings: FormSettings;
  theme: FormTheme;
  isTemplate: boolean;
  submissionCount: number;
  lastSubmissionAt?: Date;
  coverImage?: string;
}

export interface FormVersion {
  id: string;
  formId: string;
  version: number;
  createdAt: Date;
  createdBy?: string;
  structure: FormStructure;
}

export interface FormStructure {
  elements: FormElement[];
  sections: FormSection[];
  pages: FormPage[];
}

export interface FormPage {
  id: string;
  title: string;
  description?: string;
  order: number;
  sections: string[]; // Array of section IDs
  conditional?: FormElementConditional;
}

export interface FormSection {
  id: string;
  title?: string;
  description?: string;
  columns: number;
  order: number;
  elements: string[]; // Element IDs
  conditional?: FormElementConditional;
  backgroundColor?: string;
  borderRadius?: string;
  pageId?: string;
}

export interface FormElement {
  id: string;
  formId: string;
  version: number;
  type: FormElementType;
  label: string;
  required: boolean;
  properties: FormElementProperties;
  validation: FormElementValidation;
  conditional?: FormElementConditional;
  order: number;
  parentId?: string;
  sectionId?: string;
}

export type FormElementType = 
  | 'text' | 'email' | 'number' | 'textarea'
  | 'select' | 'checkbox' | 'radio' | 'date' | 'file'
  | 'section' | 'page'
  | 'rating' | 'nps' | 'matrix' | 'signature'
  | 'payment' | 'appointment' | 'address'
  | 'divider' | 'richtext' | 'image' | 'video'
  | 'calculation' | 'hidden' | 'phone' | 'url';

export interface FormElementProperties {
  placeholder?: string;
  defaultValue?: string;
  options?: string[];
  min?: number;
  max?: number;
  step?: number;
  rows?: number;
  accept?: string;
  multiple?: boolean;
  columns?: number; // For multi-column layout
  alignment?: 'left' | 'center' | 'right';
  mediaUrl?: string; // For image/video elements
  ratingScale?: number; // 5, 10, etc.
  currencyCode?: string; // For payment fields
  calculationFormula?: string; // For calculated fields
  width?: string; // For controlling element width
  buttonText?: string; // For file upload, payment buttons
  autoComplete?: boolean; // For address fields
  matrixRows?: string[];
  matrixColumns?: string[];
}

export interface FormElementValidation {
  required?: boolean;
  minLength?: number;
  maxLength?: number;
  pattern?: string;
  min?: number;
  max?: number;
  custom?: string;
  errorMessage?: string;
}

export interface FormElementConditional {
  conditions: Array<{
    field: string;
    operator: 'equals' | 'notEquals' | 'contains' | 'greaterThan' | 'lessThan' | 'isEmpty' | 'isNotEmpty';
    value: string | number | boolean;
  }>;
  logicType: 'and' | 'or';
  action: 'show' | 'hide' | 'require' | 'skip' | 'jump';
  target?: string; // Page or element ID for jump action
}

export interface FormSettings {
  allowMultipleSubmissions: boolean;
  showProgressBar: boolean;
  showPageNumbers: boolean;
  submitButtonText: string;
  successMessage: string;
  redirectUrl?: string;
  captchaEnabled: boolean;
  storeResponses: boolean;
  notificationEmails?: string[];
  autoResponseEmail?: {
    enabled: boolean;
    subject: string;
    message: string;
    includeResponses: boolean;
  };
  scoring?: {
    enabled: boolean;
    passScore?: number;
    showScoreToRespondent: boolean;
  };
}

export interface FormTheme {
  primaryColor: string;
  secondaryColor: string;
  fontFamily: string;
  fontSize: string;
  borderRadius: string;
  spacing: string;
  coverImage?: string;
  questionAnimation?: 'fade' | 'slide' | 'none';
  customCSS?: string;
  progressStyle?: 'bar' | 'dots' | 'numbers' | 'none';
  cardStyle?: 'flat' | 'shadow' | 'bordered';
  buttonStyle?: 'solid' | 'outline' | 'text';
}

export interface FormIntegration {
  id: string;
  formId: string;
  type: 'zapier' | 'webhook' | 'email' | 'google_sheets' | 'crm' | 'payment_processor';
  config: Record<string, any>;
  createdAt: Date;
  updatedAt: Date;
  lastExecutedAt?: Date;
  status: 'active' | 'inactive' | 'error';
  triggerOn: 'submission' | 'specific_answer';
  condition?: FormElementConditional;
}
```

## API Endpoints

### Authentication
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login user
- `POST /api/auth/logout` - Logout user
- `GET /api/auth/me` - Get current user info
- `POST /api/auth/forgot-password` - Request password reset
- `POST /api/auth/reset-password` - Reset password

#### Two-Factor Authentication (2FA)
- `GET /api/auth/2fa/setup` - Generate 2FA setup information (secret, QR code)
- `POST /api/auth/2fa/enable` - Enable 2FA with verification code
- `POST /api/auth/2fa/disable` - Disable 2FA with verification code
- `POST /api/auth/2fa/verify` - Verify a 2FA code during login
- `GET /api/auth/2fa/recovery-codes` - Get recovery codes
- `POST /api/auth/2fa/recovery-codes/regenerate` - Regenerate recovery codes

### Forms
- `GET /api/forms` - Get all forms for current user
- `POST /api/forms` - Create new form
- `GET /api/forms/{id}` - Get form by ID
- `PUT /api/forms/{id}` - Update form
- `DELETE /api/forms/{id}` - Delete form
- `POST /api/forms/{id}/duplicate` - Duplicate form
- `GET /api/forms/{id}/versions` - Get form versions
- `GET /api/forms/{id}/versions/{version}` - Get specific form version
- `POST /api/forms/{id}/publish` - Publish form
- `POST /api/forms/{id}/unpublish` - Unpublish form

### Form Pages
- `GET /api/forms/{formId}/pages` - Get all pages for a form
- `POST /api/forms/{formId}/pages` - Create new page
- `GET /api/forms/{formId}/pages/{id}` - Get page by ID
- `PUT /api/forms/{formId}/pages/{id}` - Update page
- `DELETE /api/forms/{formId}/pages/{id}` - Delete page
- `PUT /api/forms/{formId}/pages/reorder` - Reorder pages

### Form Sections
- `GET /api/forms/{formId}/sections` - Get all sections for a form
- `POST /api/forms/{formId}/sections` - Create new section
- `GET /api/forms/{formId}/sections/{id}` - Get section by ID
- `PUT /api/forms/{formId}/sections/{id}` - Update section
- `DELETE /api/forms/{formId}/sections/{id}` - Delete section
- `PUT /api/forms/{formId}/sections/reorder` - Reorder sections

### Form Elements
- `GET /api/forms/{formId}/elements` - Get all elements for a form
- `POST /api/forms/{formId}/elements` - Create new element
- `GET /api/forms/{formId}/elements/{id}` - Get element by ID
- `PUT /api/forms/{formId}/elements/{id}` - Update element
- `DELETE /api/forms/{formId}/elements/{id}` - Delete element
- `PUT /api/forms/{formId}/elements/reorder` - Reorder elements

### Form Submissions
- `GET /api/forms/{formId}/submissions` - Get all submissions for a form
- `POST /api/forms/{formId}/submissions` - Create new submission
- `GET /api/forms/{formId}/submissions/{id}` - Get submission by ID
- `DELETE /api/forms/{formId}/submissions/{id}` - Delete submission
- `GET /api/forms/{formId}/submissions/export` - Export submissions (CSV/Excel)

### Teams
- `GET /api/teams` - Get all teams for current user
- `POST /api/teams` - Create new team
- `GET /api/teams/{id}` - Get team by ID
- `PUT /api/teams/{id}` - Update team
- `DELETE /api/teams/{id}` - Delete team
- `GET /api/teams/{id}/members` - Get team members
- `POST /api/teams/{id}/members` - Add team member
- `DELETE /api/teams/{id}/members/{userId}` - Remove team member
- `PUT /api/teams/{id}/members/{userId}/role` - Update member role

### Form Sharing
- `GET /api/forms/{formId}/sharing` - Get sharing settings
- `POST /api/forms/{formId}/sharing` - Share form with user
- `DELETE /api/forms/{formId}/sharing/{id}` - Remove sharing
- `PUT /api/forms/{formId}/sharing/{id}` - Update sharing settings

### Form Integrations
- `GET /api/forms/{formId}/integrations` - Get all integrations
- `POST /api/forms/{formId}/integrations` - Create new integration
- `GET /api/forms/{formId}/integrations/{id}` - Get integration by ID
- `PUT /api/forms/{formId}/integrations/{id}` - Update integration
- `DELETE /api/forms/{formId}/integrations/{id}` - Delete integration
- `POST /api/forms/{formId}/integrations/{id}/test` - Test integration

### Analytics
- `GET /api/forms/{formId}/analytics` - Get form analytics
- `GET /api/forms/{formId}/analytics/submissions` - Get submission analytics
- `GET /api/forms/{formId}/analytics/views` - Get view analytics
- `GET /api/forms/{formId}/analytics/completions` - Get completion rate analytics
- `GET /api/forms/{formId}/analytics/time` - Get completion time analytics
- `GET /api/forms/{formId}/analytics/dropoff` - Get dropoff analytics

### Payments
- `GET /api/forms/{formId}/payments` - Get payment transactions for a form
- `POST /api/forms/{formId}/payments/setup` - Set up payment integration
- `GET /api/forms/{formId}/payments/{id}` - Get payment transaction details

## Frontend Component Structure

### Core Pages
- `/` - Landing page
- `/dashboard` - User dashboard
- `/forms` - Forms list
- `/forms/[id]/edit` - Form builder
- `/forms/[id]/view` - Form preview
- `/forms/[id]/submissions` - Form submissions
- `/forms/[id]/analytics` - Form analytics
- `/settings` - User settings
- `/teams` - Teams management
- `/teams/[id]` - Team details
- `/templates` - Form templates

### Form Builder Components
- `FormBuilder` - Main form builder container
- `ElementPalette` - Sidebar with available elements
- `Canvas` - Form building area
- `ElementItem` - Individual form element
- `PropertyEditor` - Element configuration panel
- `LogicBuilder` - Conditional logic editor
- `PreviewModal` - Form preview popup
- `SectionManager` - Manage form sections and layout
- `PageManager` - Manage multi-page forms
- `SettingsPanel` - Form settings configuration panel
- `ThemeCustomizer` - Form theme and styling customization
- `IntegrationManager` - Configure third-party integrations

### Dashboard Components
- `FormsList` - List of user's forms
- `FormCard` - Individual form card
- `AnalyticsSummary` - Overview of form performance
- `RecentSubmissions` - Recent form submissions
- `FormFilters` - Filter and search forms
- `TeamAccess` - Team collaboration management

### Form View Components
- `FormRenderer` - Renders the form for submission
- `FormElement` - Renders individual form elements
- `ProgressBar` - Shows multi-page form progress
- `ValidationErrors` - Displays form validation errors
- `SubmissionSuccess` - Success page after form submission
- `PaymentProcessor` - Handles payment form elements
- `ConditionalDisplay` - Handles showing/hiding elements based on logic

## Form Creation Approach

Based on market research and user pain points analysis, FormFusion will implement a hybrid approach to form creation that combines multiple methods:

### 1. Primary Interface: Enhanced Drag and Drop

- **Modern Visual Builder**: Implement an intuitive drag and drop interface as the primary form creation method
- **Fluid Interactions**: Use Aceternity UI components to create smooth, physics-based animations for drag and drop operations
- **Grid/Section System**: Provide a robust layout system with customizable columns and sections
- **Visual Cues**: Clear drop zones, alignment guides, and visual feedback during element manipulation
- **Contextual Controls**: Element-specific controls that appear on hover/selection
- **Responsive Preview**: Real-time responsive design preview across device sizes

### 2. Secondary Interface: Text-Based Quick Entry

- **Text Editor Mode**: Provide an alternative text editor style interface similar to Tally.so
- **Mode Switching**: Allow seamless switching between visual and text modes
- **Keyboard Shortcuts**: Extensive keyboard shortcuts for power users
- **Quick Commands**: Support for slash commands to quickly add and configure elements
- **Markdown Support**: Allow markdown syntax for formatting descriptions and labels
- **Bulk Editing**: Support for editing multiple elements simultaneously in text mode

### 3. Template System with AI Enhancement

- **Smart Templates**: AI-powered suggestions based on form purpose and industry
- **Template Library**: Extensive collection of pre-built templates for common scenarios
- **Template Customization**: Modify templates through both visual and text interfaces
- **Sectional Templates**: Reusable section templates that can be added to any form
- **AI Optimization**: Intelligent suggestions for improving form layouts and questions
- **Industry-Specific Templates**: Specialized templates for different verticals (healthcare, education, etc.)

### 4. Collaborative Creation

- **Real-time Co-editing**: Multiple team members can work on a form simultaneously
- **Comments and Suggestions**: Leave feedback directly on form elements
- **Version History**: Track changes and restore previous versions
- **Role-based Access**: Different permissions for different team members
- **Approval Workflows**: Multi-step approval process for form publication

### Benefits of This Hybrid Approach

- **Addresses diverse user preferences**: Different users can work in their preferred style
- **Balances simplicity and power**: Easy for beginners but powerful enough for experts
- **Improves efficiency**: Provides the most efficient interface for each task
- **Enhances collaboration**: Supports team-based form development workflows
- **Leverages AI capabilities**: Uses artificial intelligence to improve form quality

### Implementation Considerations

- **Consistent Data Model**: Regardless of creation method, all interfaces must update the same underlying form model
- **Synchronized Views**: Changes in one interface should immediately reflect in others
- **Performance Optimization**: Implement virtualization for large forms and selective rendering
- **Accessibility Support**: Ensure keyboard navigation works well for the drag and drop interface
- **Progressive Enhancement**: Core functionality should work without advanced features

This hybrid approach directly addresses key pain points identified in market research:
- **Design limitations**: Solved with better layout controls and visual customization
- **Complex interfaces**: Mitigated by offering multiple entry methods for different user preferences
- **Technical frustrations**: Simplified by making complex logic creation more intuitive

## Development Status

### Completed Features


### In Progress


### Pending Features


## Implementation Roadmap

### Phase 1: Core Platform (MVP)
- Basic authentication and user management
- Form builder with essential elements
- Form submission and data collection
- Form management dashboard
- Basic analytics

### Phase 2: Enhanced Form Building
- Section layout management
- Advanced element types
- Rich text editing
- Enhanced theme customization
- Form templates library

### Phase 3: Advanced Features
- Conditional logic and workflows
- Payment integration
- File uploads and management
- Form templates and duplicating
- Advanced validation rules

### Phase 4: AI and Optimization
- AI-powered form suggestions
- Completion rate optimization
- A/B testing capabilities
- User behavior analytics
- Form scoring and calculations

### Phase 5: Integration Ecosystem
- API for external access
- Third-party integrations (CRM, email, etc.)
- Webhooks and automation
- Data export and reporting
- Custom API endpoints for clients

### Phase 6: Enterprise & Compliance
- Enterprise SSO integration
- Advanced spam protection
- GDPR/data privacy compliance tools
- Advanced security features

### Phase 7: Global & Accessibility
- Localization/internationalization support
- Accessibility compliance (WCAG standards)
- Offline form submission support
- Data migration tools

### Phase 8: Advanced Customization
- Enhanced white labeling capabilities
- Custom notification workflows
- Advanced theming options
- Custom branding control

## Development Guidelines

### Coding Standards
- Use TypeScript for all frontend code
- Follow C# coding conventions for backend
- Use async/await for asynchronous operations
- Write unit tests for all business logic
- Document all public APIs and components

### Git Workflow
- Main branch for production
- Development branch for ongoing work
- Feature branches for new features
- Pull requests with code reviews

### CI/CD Pipeline
- Automated testing on pull requests
- Build verification
- Staging deployment
- Production deployment

### Testing Strategy
- Unit tests for business logic
- Integration tests for API endpoints
- E2E tests for critical user flows
- Performance testing for form submissions
- Accessibility testing for form rendering

## Security Considerations

- HTTPS for all communications
- JWT for authentication
- Two-factor authentication (2FA)
- Input validation for all form submissions
- CSRF protection
- Rate limiting for form submissions
- Data encryption for sensitive information
- Regular security audits

### Two-Factor Authentication Implementation

The two-factor authentication system will use Time-based One-Time Password (TOTP) protocol compatible with authenticator apps like Google Authenticator, Microsoft Authenticator, and Authy. Key implementation details:

1. **Setup Process**:
   - Generate a cryptographically secure random secret for each user
   - Display QR code containing the otpauth:// URL
   - Verify initial code before enabling 2FA
   - Generate and store recovery codes (one-time use)

2. **Authentication Flow**:
   - Standard login with username/password
   - If 2FA enabled, prompt for verification code
   - Validate TOTP code (with small time window allowance)
   - Alternative recovery code path for emergency access

3. **Security Measures**:
   - Rate limiting on verification attempts
   - Secret stored using encryption
   - Recovery codes stored as hashes
   - Audit logging for all 2FA-related actions

4. **User Experience**:
   - Clear setup instructions
   - Backup code download/copy
   - Device management
   - Account recovery options

## Performance Optimizations

- Redis caching for frequently accessed forms
- CDN for static assets
- Lazy loading for form builder components
- Pagination for form submissions
- Optimized database queries with proper indexing
- Image and file optimization

## Monitoring and Operations

- Application performance monitoring
- Error tracking and reporting
- Usage analytics
- Server health monitoring
- Backup and disaster recovery
- Scalability planning

---

This technical specification will be updated as development progresses. Features will be marked as completed when they are fully implemented and tested.