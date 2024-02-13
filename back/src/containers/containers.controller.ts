import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
} from '@nestjs/common';
import { ContainersService } from './containers.service';
import { CreateContainerDto } from './dto/create-container.dto';
import { UpdateContainerDto } from './dto/update-container.dto';
import { FeatureFlagsService } from 'src/feature-flags/feature-flags.service';

@Controller('containers')
export class ContainersController {
  constructor(
    private readonly containersService: ContainersService,
    private readonly featureFlagsService: FeatureFlagsService
  ) {}

  @Post()
  create(@Body() createContainerDto: CreateContainerDto) {
    return this.containersService.create(createContainerDto);
  }

  @Get()
  findAll() {
    return this.containersService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.containersService.findOne(+id);
  }

  @Get(':id/feature-flags')
  findFeatureFlags(@Param('id') id: string) {
    return this.featureFlagsService.findFeatureFlagsByContainerId(+id);
  }

  @Patch(':id')
  update(
    @Param('id') id: string,
    @Body() updateContainerDto: UpdateContainerDto
  ) {
    return this.containersService.update(+id, updateContainerDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.containersService.remove(+id);
  }
}
