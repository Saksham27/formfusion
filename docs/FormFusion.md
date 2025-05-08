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
    CONSTRAINT FK_FormElements_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT FK_FormElements_Parent FOREIGN KEY (ParentId) REFERENCES FormElements(Id)
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
    CONSTRAINT FK_FormAnalytics_Forms FOREIGN KEY (FormId) REFERENCES Forms(Id),
    CONSTRAINT UQ_FormAnalytics_FormDate UNIQUE (FormId, Date)
);
```

## Entity Relationships

### Key Relationships
- User can own multiple Forms
- Team can own multiple Forms
- Form has multiple FormElements
- Form has multiple FormSubmissions
- FormSubmission has multiple SubmissionValues
- User can be part of multiple Teams (through TeamMembers)
- Form can be shared with multiple Users (through FormSharing)
- Form can have multiple Integrations (through FormIntegrations)

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

### Dashboard Components
- `FormsList` - List of user's forms
- `FormCard` - Individual form card
- `AnalyticsSummary` - Overview of form performance
- `RecentSubmissions` - Recent form submissions
- `FormFilters` - Filter and search forms

### Form View Components
- `FormRenderer` - Renders the form for submission
- `FormElement` - Renders individual form elements
- `ProgressBar` - Shows multi-page form progress
- `ValidationErrors` - Displays form validation errors
- `SubmissionSuccess` - Success page after form submission

## Development Status

### Planned Features
- [ ] Project structure and setup
- [ ] Database schema design
- [ ] User authentication (backend)
- [ ] User authentication (frontend)
- [ ] Form CRUD operations (backend)
- [ ] Form CRUD operations (frontend)
- [ ] Form builder UI components
- [ ] Element drag-and-drop functionality
- [ ] Property editor for elements
- [ ] Form renderer for submissions
- [ ] Form submission handling
- [ ] Form analytics (basic)
- [ ] Team management
- [ ] User settings
- [ ] Form sharing
- [ ] Conditional logic builder
- [ ] Form versioning
- [ ] Form templates
- [ ] Submission exports

### Phase 1: Core Platform Features
- [ ] AI-powered form optimization
- [ ] Real-time collaboration
- [ ] Advanced analytics
- [ ] Integration ecosystem
- [ ] Mobile app for form management
- [ ] PDF generation and document workflows

### Phase 2: Extended Features
- [ ] Localization/internationalization support for multilingual forms
- [ ] Accessibility compliance features (WCAG standards)
- [ ] Custom themes/white labeling
- [ ] Offline form submission support
- [ ] Data migration tools for users coming from other platforms
- [ ] Enterprise SSO integration
- [ ] Custom notification workflows for form submissions
- [ ] Advanced spam protection
- [ ] GDPR/data privacy compliance tools
- [ ] Custom API endpoints for client embedding

## Implementation Roadmap

### Phase 1: Core Platform (MVP)
- Basic authentication and user management
- Form builder with essential elements
- Form submission and data collection
- Form management dashboard
- Basic analytics

### Phase 2: Team Collaboration
- Team management
- Form sharing and permissions
- Commenting and feedback
- Version history and rollback

### Phase 3: Advanced Features
- Conditional logic and workflows
- Payment integration
- File uploads and management
- Form templates and duplicating

### Phase 4: AI and Optimization
- AI-powered form suggestions
- Completion rate optimization
- A/B testing capabilities
- User behavior analytics

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

## Database Evolution Plan

The existing database schema is designed to support progressive enhancement without major restructuring. Future features will be implemented through:

### Phase 5-6 Additions
```sql
-- For Custom API endpoints
CREATE TABLE ApiKeys (
    Id UUID PRIMARY KEY,
    UserId UUID NOT NULL REFERENCES Users(Id),
    FormId UUID REFERENCES Forms(Id), -- NULL means all forms
    Name VARCHAR(255) NOT NULL,
    Key VARCHAR(255) NOT NULL UNIQUE,
    Permissions JSONB NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastUsedAt TIMESTAMP
);

-- For Enterprise SSO
ALTER TABLE Users ADD COLUMN SsoProvider VARCHAR(50);
ALTER TABLE Users ADD COLUMN SsoIdentifier VARCHAR(255);

CREATE TABLE SsoConfigurations (
    Id UUID PRIMARY KEY,
    TeamId UUID NOT NULL REFERENCES Teams(Id),
    Provider VARCHAR(50) NOT NULL,
    Configuration JSONB NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- For GDPR compliance
ALTER TABLE Forms ADD COLUMN DataRetentionDays INTEGER;
ALTER TABLE FormElements ADD COLUMN ContainsPII BOOLEAN DEFAULT FALSE;

CREATE TABLE DataDeletionRequests (
    Id UUID PRIMARY KEY,
    SubmissionId UUID REFERENCES FormSubmissions(Id),
    RequestedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Status VARCHAR(50) DEFAULT 'pending',
    CompletedAt TIMESTAMP,
    RequestorIdentifier VARCHAR(255) NOT NULL
);
```

### Phase 7-8 Additions
```sql
-- For Localization support
CREATE TABLE ElementTranslations (
    Id UUID PRIMARY KEY,
    ElementId UUID NOT NULL REFERENCES FormElements(Id),
    LanguageCode VARCHAR(10) NOT NULL,
    TranslatedLabel TEXT,
    TranslatedHelpText TEXT,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT UQ_ElementTranslations UNIQUE (ElementId, LanguageCode)
);

ALTER TABLE Forms ADD COLUMN DefaultLanguage VARCHAR(10) DEFAULT 'en';
ALTER TABLE Forms ADD COLUMN SupportedLanguages JSONB;

-- For Custom notification workflows
CREATE TABLE NotificationWorkflows (
    Id UUID PRIMARY KEY,
    FormId UUID NOT NULL REFERENCES Forms(Id),
    Name VARCHAR(255) NOT NULL,
    Triggers JSONB NOT NULL, -- Event types that trigger this workflow
    Actions JSONB NOT NULL, -- Actions to take (email, webhook, etc.)
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

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